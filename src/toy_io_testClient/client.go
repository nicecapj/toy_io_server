package main

import (
	Network "Network"
	"log"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
	"util"

	"github.com/golang/protobuf/proto"
)

// User is session + logic packet
type ClientSession struct {
	*Network.Session
}

// Init used for initialize of session
func (this *ClientSession) Init(session *Network.Session) {
	this.Session = session
}

// DispatchPacket is dispatch packet.
func (this *ClientSession) DispatchPacket(protocolID PROTOCOL.ProtocolID, buffer []byte) {
	switch protocolID {
	case PROTOCOL.ProtocolID_LoginRes:
		{
			res := &LobbyPacket.LoginRes{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)
			log.Printf("%s\n", res.String())
		}
	}
}

//-----------------------------------------------------------------------------
func (this *ClientSession) reqestLogin(userName string) {

	req := &LobbyPacket.LoginReq{}
	req.Name = userName
	this.SendPacket(PROTOCOL.ProtocolID_LoginReq, req)
}
