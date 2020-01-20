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

type PacketHandler func(buffer []byte)

// User is session + logic packet
type User struct {
	*Network.Session
	enteredRoom     *Room
	isReadyGame     bool
	currentLocation packet_lobby.Location
	targetLocation  packet_lobby.Location

	handle map[PROTOCOL.ProtocolID]PacketHandler
}

// Init used for initialize of session
func (user *User) Init(session *Network.Session) {
	user.Session = session
	user.InitHandler()

	user.currentLocation.X = 0
	user.currentLocation.Y = 0
	user.currentLocation.Z = 0
	user.targetLocation.X = 0
	user.targetLocation.Y = 0
	user.targetLocation.Z = 0
}

func (user *User) InitHandler() {
	user.handle = make(map[PROTOCOL.ProtocolID]PacketHandler)

	user.handle[PROTOCOL.ProtocolID_LoginReq] = user.HandleLoginReq
	user.handle[PROTOCOL.ProtocolID_RoomEnterReq] = user.HandleRoomEnterReq
	user.handle[PROTOCOL.ProtocolID_ReadyForGameReq] = user.HandleReadyForGameReq
	user.handle[PROTOCOL.ProtocolID_RoomLeaveReq] = user.HandleRoomLeaveReq
	user.handle[PROTOCOL.ProtocolID_MoveStartReq] = user.HandleMoveStartReq
	user.handle[PROTOCOL.ProtocolID_MoveChangeReq] = user.HandleMoveChangeReq
	user.handle[PROTOCOL.ProtocolID_MoveEndReq] = user.HandleMoveEndReq
	user.handle[PROTOCOL.ProtocolID_FisingStartReq] = user.HandleFisingStartReq
}

// DispatchPacket is dispatch packet.
func (user *User) HandlePacket(protocolID PROTOCOL.ProtocolID, buffer []byte) {
	user.handle[protocolID](buffer)
}

func (user *User) Close() {
	if user.enteredRoom != nil {
		user.enteredRoom.Leave(user)
	}

	user.Session.Close()
}

//=================================================================================================
//PACKET HANDLERS
//=================================================================================================
func (user *User) HandleLoginReq(buffer []byte) {
	req := &LobbyPacket.LoginReq{}
	err := proto.Unmarshal(buffer[:], req)
	Util.ProcessError(err)
	log.Printf("%s\n", req.String())

	user.OnLoginReq(req)
}

func (user *User) HandleRoomEnterReq(buffer []byte) {
	req := &LobbyPacket.RoomEnterReq{}
	err := proto.Unmarshal(buffer[:], req)
	Util.ProcessError(err)
	log.Printf("%s\n", req.String())

	user.OnRoomEnterReq(req)
}

func (user *User) HandleReadyForGameReq(buffer []byte) {
	req := &LobbyPacket.ReadyForGameReq{}
	err := proto.Unmarshal(buffer[:], req)
	Util.ProcessError(err)
	log.Printf("%s\n", req.String())

	user.OnReadyForGameReq(req)
}

func (user *User) HandleRoomLeaveReq(buffer []byte) {
	req := &LobbyPacket.RoomLeaveReq{}
	err := proto.Unmarshal(buffer[:], req)
	Util.ProcessError(err)
	log.Printf("%s\n", req.String())

	user.OnRoomLeaveReq(req)
}

func (user *User) HandleMoveStartReq(buffer []byte) {
	req := &LobbyPacket.MoveStartReq{}
	err := proto.Unmarshal(buffer[:], req)
	Util.ProcessError(err)
	log.Printf("%s\n", req.String())

	user.OnMoveStartReq(req)
}

func (user *User) HandleMoveChangeReq(buffer []byte) {
	req := &LobbyPacket.MoveChangeReq{}
	err := proto.Unmarshal(buffer[:], req)
	Util.ProcessError(err)
	log.Printf("%s\n", req.String())

	user.OnMoveChangeReq(req)
}

func (user *User) HandleMoveEndReq(buffer []byte) {
	req := &LobbyPacket.MoveEndReq{}
	err := proto.Unmarshal(buffer[:], req)
	Util.ProcessError(err)
	log.Printf("%s\n", req.String())

	user.OnMoveEndReq(req)
}

func (user *User) HandleFisingStartReq(buffer []byte) {
	{
		req := &LobbyPacket.FisingStartReq{}
		err := proto.Unmarshal(buffer[:], req)
		Util.ProcessError(err)
		log.Printf("%s\n", req.String())

		user.OnFisingStartReq(req)
	}
}

//=================================================================================================
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

func (user *User) OnFisingStartReq(req *LobbyPacket.FisingStartReq) {

	//var fishId int32
	//var cm int32

	fishingManager := GetFishingManager()
	fishId, cm := fishingManager.CatchFish(req.AreaId)

	timer := time.NewTimer(2000 * time.Millisecond)
	go func() {
		<-timer.C

		hitNfy := &LobbyPacket.FisingHitNfy{}
		hitNfy.CasterUid = user.UID
		hitNfy.FishId = fishId
		hitNfy.FishCM = cm
		user.enteredRoom.Broadcast(user, PROTOCOL.ProtocolID_FisingHitNfy, hitNfy)
	}()

	res := &LobbyPacket.FisingStartRes{}
	res.RetCode = ReturnCode.ReturnCode_retOK
	user.SendPacket(PROTOCOL.ProtocolID_FisingStartRes, res)

	nfy := &LobbyPacket.FisingStartNfy{}
	nfy.CasterUid = user.UID
	nfy.RodType = 1
	user.enteredRoom.Broadcast(user, PROTOCOL.ProtocolID_FisingStartNfy, nfy)
}

func (user *User) RandomMove(duration time.Duration) {
	timer := time.NewTimer(duration)
	func() {
		<-timer.C

		rand.NewSource(time.Now().UnixNano())

		if rand.Int31n(10) > 5 {
			user.targetLocation.X += rand.Int31n(3)
			user.targetLocation.Y += rand.Int31n(3)
		} else {
			user.targetLocation.X -= rand.Int31n(3)
			user.targetLocation.Y -= rand.Int31n(3)
		}
	}()
}
