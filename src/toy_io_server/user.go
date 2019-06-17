package main

import (
	Network "Network"
	"log"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
	ReturnCode "packet_returncode"
	Util "util"

	"github.com/golang/protobuf/proto"
)

// User is session + logic packet
type User struct {
	*Network.Session
	enteredRoom *Room
	isReadyGame bool
}

// Init used for initialize of session
func (user *User) Init(session *Network.Session) {
	user.Session = session
}

func (user *User) Close() {
	if user.enteredRoom != nil {
		user.enteredRoom.Leave(user)
	}

	user.Session.Close()
}

// DispatchPacket is dispatch packet.
func (user *User) DispatchPacket(protocolID PROTOCOL.ProtocolID, buffer []byte) {
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
