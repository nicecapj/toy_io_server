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
}

// Init used for initialize of session
func (user *User) Init(session *Network.Session) {
	user.Session = session
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
	res := &LobbyPacket.RoomEnterRes{}

	roomManager := GetRoomManager()
	room := roomManager.GetLeisuerlyRoom()
	if room != nil {
		if room.Enter(user) == true {
			res.RetCode = ReturnCode.ReturnCode_retOK
			res.RoomID = room.RoomID
			res.RoomName = room.Name
		} else {
			//fail to enter room
			res.RetCode = ReturnCode.ReturnCode_retFail
		}

	} else {
		//not enought room.
		res.RetCode = ReturnCode.ReturnCode_retFail
	}

	user.SendPacket(PROTOCOL.ProtocolID_RoomEnterRes, res)
}
