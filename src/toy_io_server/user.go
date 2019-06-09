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

func (user *User) HandlePacket(bufferArray []byte) {
	//부모안에서 처리후에, 자식의 함수로 virtual처럼 처리하고 싶은데 못했음. 방법 찾는중
	//user.Session.HandlePacket(bufferArray)

	//header빼오는 부분은 Session으로 감추고 싶다...
	header, err := Network.GetHeader(bufferArray[:Network.MaxPacketSize])
	if err != nil {
		log.Panicln("read header")
	}

	user.OnReceived(header.PacketID, bufferArray, header.PacketSize)
}

// OnReceived dispatch packet.
//func (session *Network.Session) OnReceived(protocolID PROTOCOL.ProtocolID, buffer []byte) {
func (user *User) OnReceived(protocolID PROTOCOL.ProtocolID, buffer []byte, packetSize int32) {
	switch protocolID {
	case PROTOCOL.ProtocolID_LoginReq:
		{
			req := &LobbyPacket.LoginReq{}
			err := proto.Unmarshal(buffer[Network.PacketHeaderLen:packetSize], req)
			Util.ProcessError(err)
			log.Printf("%s\n", req.String())

			user.OnLoginReq(req)
		}
	case PROTOCOL.ProtocolID_LoginRes:
		{
			res := &LobbyPacket.LoginRes{}
			err := proto.Unmarshal(buffer[Network.PacketHeaderLen:packetSize], res)
			Util.ProcessError(err)
			log.Printf("%s\n", res.String())
		}
	}
}

// OnLoginReq is handler for login request from client
func (user *User) OnLoginReq(req *LobbyPacket.LoginReq) {
	sessionManager := Network.GetSessionManager()

	name := req.GetName()
	user.Name = name

	res := &LobbyPacket.LoginRes{}
	if sessionManager.FindUser(name) {
		res.RetCode = ReturnCode.ReturnCode_retExist
		res.Uid = sessionManager.GetUID(name)
	} else {
		sessionManager.AddUser(name)
		res.RetCode = ReturnCode.ReturnCode_retOK
		res.Uid = sessionManager.GetUID(name)
	}
	user.SendPacket(PROTOCOL.ProtocolID_LoginRes, res)
}
