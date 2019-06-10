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
func (this *User) Init(session *Network.Session) {
	this.Session = session
}

// DispatchPacket is dispatch packet.
func (this *User) DispatchPacket(protocolID PROTOCOL.ProtocolID, buffer []byte) {
	switch protocolID {
	case PROTOCOL.ProtocolID_LoginReq:
		{
			req := &LobbyPacket.LoginReq{}
			err := proto.Unmarshal(buffer[:], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			this.OnLoginReq(req)
		}
	}
}

// OnLoginReq is handler for login request from client
func (this *User) OnLoginReq(req *LobbyPacket.LoginReq) {
	sessionManager := Network.GetSessionManager()

	name := req.GetName()
	this.Name = name

	res := &LobbyPacket.LoginRes{}
	if sessionManager.FindUser(name) {
		res.RetCode = ReturnCode.ReturnCode_retExist
		res.Uid = sessionManager.GetUID(name)
	} else {
		sessionManager.AddUser(name)
		res.RetCode = ReturnCode.ReturnCode_retOK
		res.Uid = sessionManager.GetUID(name)
	}
	this.SendPacket(PROTOCOL.ProtocolID_LoginRes, res)
}
