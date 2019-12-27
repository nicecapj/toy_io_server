package main

import (
	Network "Network"
	"log"
	"math/rand"
	"packet_lobby"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
	ReturnCode "packet_returncode"
	"time"
	Util "util"

	"github.com/golang/protobuf/proto"
)

// User is session + logic packet
type User struct {
	*Network.Session
	enteredRoom     *Room
	isReadyGame     bool
	currentLocation packet_lobby.Location
	targetLocation  packet_lobby.Location
}

// Init used for initialize of session
func (user *User) Init(session *Network.Session) {
	user.Session = session

	user.currentLocation.X = 0
	user.currentLocation.Y = 0
	user.currentLocation.Z = 0
	user.targetLocation.X = 0
	user.targetLocation.Y = 0
	user.targetLocation.Z = 0
}

func (user *User) Close() {
	if user.enteredRoom != nil {
		user.enteredRoom.Leave(user)
	}

	user.Session.Close()
}

// DispatchPacket is dispatch packet.
func (user *User) HandlePacket(protocolID PROTOCOL.ProtocolID, buffer []byte) {

	switch protocolID {
	case PROTOCOL.ProtocolID_LoginReq:
		{
			req := &LobbyPacket.LoginReq{}
			err := proto.Unmarshal(buffer[:], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			user.OnLoginReq(req)
		}

	case PROTOCOL.ProtocolID_RoomEnterReq:
		{
			req := &LobbyPacket.RoomEnterReq{}
			err := proto.Unmarshal(buffer[:], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			user.OnRoomEnterReq(req)
		}
	case PROTOCOL.ProtocolID_ReadyForGameReq:
		{
			req := &LobbyPacket.ReadyForGameReq{}
			err := proto.Unmarshal(buffer[:], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			user.OnReadyForGameReq(req)
		}

	case PROTOCOL.ProtocolID_RoomLeaveReq:
		{
			req := &LobbyPacket.RoomLeaveReq{}
			err := proto.Unmarshal(buffer[:], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			user.OnRoomLeaveReq(req)
		}

	case PROTOCOL.ProtocolID_MoveStartReq:
		{
			req := &LobbyPacket.MoveStartReq{}
			err := proto.Unmarshal(buffer[:], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			user.OnMoveStartReq(req)
		}

	case PROTOCOL.ProtocolID_MoveChangeReq:
		{
			req := &LobbyPacket.MoveChangeReq{}
			err := proto.Unmarshal(buffer[:], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			user.OnMoveChangeReq(req)
		}

	case PROTOCOL.ProtocolID_MoveEndReq:
		{
			req := &LobbyPacket.MoveEndReq{}
			err := proto.Unmarshal(buffer[:], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			user.OnMoveEndReq(req)
		}
	}
}

// OnLoginReq is handler for login request from client
func (user *User) OnLoginReq(req *LobbyPacket.LoginReq) {
	accountManager := Network.GetAccountManager()

	name := req.GetName()
	user.Name = name

	res := &LobbyPacket.LoginRes{}
	if accountManager.FindUser(name) {
		res.RetCode = ReturnCode.ReturnCode_retExist
	} else {
		accountManager.AddUser(name)
		res.RetCode = ReturnCode.ReturnCode_retOK
	}

	user.UID = accountManager.GetUID(name)
	res.Uid = user.UID

	user.SendPacket(PROTOCOL.ProtocolID_LoginRes, res)
}

// OnRoomEnterReq ...
func (user *User) OnRoomEnterReq(req *LobbyPacket.RoomEnterReq) {
	roomManager := GetRoomManager()
	room := roomManager.GetLeisuerlyRoom()
	if room != nil {
		if room.Enter(user) == true {
			user.SetRoom(room)
		} else {
			//error packet : fail to enter room
			//res.RetCode = ReturnCode.ReturnCode_retFail
		}

	} else {
		//error packet : not enought room.
		//res.RetCode = ReturnCode.ReturnCode_retFail
	}
}

// OnRoomLeaveReq ...
func (user *User) OnRoomLeaveReq(req *LobbyPacket.RoomLeaveReq) {
	res := &LobbyPacket.RoomLeaveRes{}

	roomManager := GetRoomManager()
	room := roomManager.FindRoom(req.RoomID)
	if room != nil {
		room.Leave(user)
	} else {
		//error packet : can`t not find room.
		res.RetCode = ReturnCode.ReturnCode_retFail
	}
}

// OnReadyForGameReq. if do not received from user, ignored from all logic
func (user *User) OnReadyForGameReq(req *LobbyPacket.ReadyForGameReq) {

	user.isReadyGame = true
}

func (user *User) SetRoom(room *Room) {
	user.Lock()
	user.enteredRoom = room
	user.Unlock()
}

func (user *User) OnMoveStartReq(req *LobbyPacket.MoveStartReq) {
	req.GetTargetPos()

	user.targetLocation.X = req.GetTargetPos().GetX()
	user.targetLocation.Y = req.GetTargetPos().GetY()

	res := &LobbyPacket.MoveStartRes{}
	res.RetCode = ReturnCode.ReturnCode_retOK
	//res.TargetPos = &LobbyPacket.Location{}
	res.TargetPos = req.TargetPos
	user.SendPacket(PROTOCOL.ProtocolID_MoveStartRes, res)

	nfy := &LobbyPacket.MoveStartNfy{}
	nfy.Uid = req.Uid
	nfy.TargetPos = &user.targetLocation
	//nfy.TargetPos.X = user.targetLocation.X
	//nfy.TargetPos.Y = user.targetLocation.Y
	//nfy.TargetPos.Z = user.targetLocation.Z

	user.enteredRoom.Broadcast(user, PROTOCOL.ProtocolID_MoveStartNfy, nfy)
}

func (user *User) OnMoveChangeReq(req *LobbyPacket.MoveChangeReq) {
	user.targetLocation.X = req.GetTargetPos().GetX()
	user.targetLocation.Y = req.GetTargetPos().GetY()

	res := &LobbyPacket.MoveChangeReq{}
	//res.RetCode = ReturnCode.ReturnCode_retOK
	user.SendPacket(PROTOCOL.ProtocolID_MoveChangeRes, res)

	nfy := &LobbyPacket.MoveChangeNfy{}
	nfy.Uid = req.Uid
	nfy.TargetPos = &user.targetLocation
	user.enteredRoom.Broadcast(user, PROTOCOL.ProtocolID_MoveChangeNfy, nfy)
}

func (user *User) OnMoveEndReq(req *LobbyPacket.MoveEndReq) {
	user.targetLocation.X = req.GetTargetPos().GetX()
	user.targetLocation.Y = req.GetTargetPos().GetY()

	res := &LobbyPacket.MoveEndRes{}
	res.RetCode = ReturnCode.ReturnCode_retOK
	user.SendPacket(PROTOCOL.ProtocolID_MoveEndRes, res)

	nfy := &LobbyPacket.MoveEndNfy{}
	nfy.Uid = req.Uid
	nfy.TargetPos = &user.targetLocation
	user.enteredRoom.Broadcast(user, PROTOCOL.ProtocolID_MoveEndNfy, nfy)

}

func (user *User) RandomMove(duration time.Duration) {
	timer := time.NewTimer(duration)
	func() {
		<-timer.C

		rand.NewSource(time.Now().UnixNano())

		user.targetLocation.X += rand.Int31n(3)
		user.targetLocation.Y += rand.Int31n(3)
	}()
}
