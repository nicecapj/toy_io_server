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

			log.Print(reflect.TypeOf(res))
			log.Printf("%s\n", res.String())

			if res.RetCode != RETURNCODE.ReturnCode_retFail {
				this.Uid = res.Uid
				this.RequestRoomEnterReq(this.Uid)
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
				this.RequestReadyForGameReq(this.Uid) //Considered to be loaded
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

//-----------------------------------------------------------------------------
func (this *ClientSession) RequestLoginReq(userName string) {

	req := &LobbyPacket.LoginReq{Name: userName}
	this.SendPacket(PROTOCOL.ProtocolID_LoginReq, req)
}

func (this *ClientSession) RequestRoomEnterReq(uid int64) {

	req := &LobbyPacket.RoomEnterReq{Uid: uid}
	this.SendPacket(PROTOCOL.ProtocolID_RoomEnterReq, req)
}

func (this *ClientSession) RequestReadyForGameReq(uid int64) {

	req := &LobbyPacket.ReadyForGameReq{Uid: uid}
	this.SendPacket(PROTOCOL.ProtocolID_ReadyForGameReq, req)
}
func (this *ClientSession) RequestRoomLeaveReq(uid int64) {

	req := &LobbyPacket.RoomLeaveReq{Uid: uid}
	this.SendPacket(PROTOCOL.ProtocolID_RoomLeaveRes, req)
}
