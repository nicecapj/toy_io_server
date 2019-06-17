package main

import (
	Network "Network"
	"log"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
	RETURNCODE "packet_returncode"
	"reflect"
	"util"

	"github.com/golang/protobuf/proto"
)

// ClientSession is session + logic packet
type ClientSession struct {
	*Network.Session
}

// Init used for initialize of session
func (clientSession *ClientSession) Init(session *Network.Session) {
	clientSession.Session = session
}

// DispatchPacket is dispatch packet.
func (clientSession *ClientSession) DispatchPacket(protocolID PROTOCOL.ProtocolID, buffer []byte) {
	switch protocolID {
	case PROTOCOL.ProtocolID_LoginRes:
		{
			res := &LobbyPacket.LoginRes{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)

			log.Print(reflect.TypeOf(res))
			log.Printf("%s\n", res.String())

			if res.RetCode != RETURNCODE.ReturnCode_retFail {
				clientSession.UID = res.Uid
				clientSession.RequestRoomEnterReq(clientSession.UID)
			}
		}
	case PROTOCOL.ProtocolID_RoomEnterRes:
		{
			res := &LobbyPacket.RoomEnterRes{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)

			log.Print(reflect.TypeOf(res))
			log.Printf("%s\n", res.String())

			if res.RetCode != RETURNCODE.ReturnCode_retFail {
				clientSession.RequestReadyForGameReq(clientSession.UID) //Considered to be loaded
			}
		}
	case PROTOCOL.ProtocolID_RoomEnterNfy:
		{
			res := &LobbyPacket.RoomEnterNfy{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)

			log.Print(reflect.TypeOf(res))
			log.Printf("%s\n", res.String())

			userList := res.GetUserInfoList()
			userCount := len(userList)
			log.Printf("userCount : %d\n", userCount)
		}
	}
}

// RequestLoginReq is ...
func (clientSession *ClientSession) RequestLoginReq(userName string) {

	req := &LobbyPacket.LoginReq{Name: userName}
	clientSession.SendPacket(PROTOCOL.ProtocolID_LoginReq, req)
}

// RequestRoomEnterReq ...
func (clientSession *ClientSession) RequestRoomEnterReq(uid int64) {

	req := &LobbyPacket.RoomEnterReq{Uid: uid}
	clientSession.SendPacket(PROTOCOL.ProtocolID_RoomEnterReq, req)
}

// RequestReadyForGameReq will request to server after client ready to play and receive packet from server(when client loaded c omplete map and game dates)
func (clientSession *ClientSession) RequestReadyForGameReq(uid int64) {

	req := &LobbyPacket.ReadyForGameReq{Uid: uid}
	clientSession.SendPacket(PROTOCOL.ProtocolID_ReadyForGameReq, req)
}

// RequestRoomLeaveReq ...
func (clientSession *ClientSession) RequestRoomLeaveReq(uid int64) {

	req := &LobbyPacket.RoomLeaveReq{Uid: uid}
	clientSession.SendPacket(PROTOCOL.ProtocolID_RoomLeaveRes, req)
}
