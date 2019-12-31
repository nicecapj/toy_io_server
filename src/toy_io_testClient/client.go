package main

import (
	Network "Network"
	"log"
	"math/rand"
	"packet_lobby"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
	RETURNCODE "packet_returncode"
	"reflect"
	"time"
	"util"

	"github.com/golang/protobuf/proto"
)

// ClientSession is session + logic packet
type ClientSession struct {
	*Network.Session

	//test
	timer        *time.Timer
	moveEndtimer *time.Timer

	roomID          int32
	roomName        string
	currentLocation packet_lobby.Location
	targetLocation  packet_lobby.Location
}

// Init used for initialize of session
func (clientSession *ClientSession) Init(session *Network.Session) {
	clientSession.Session = session

	clientSession.currentLocation.X = 0
	clientSession.currentLocation.Y = 0
	clientSession.currentLocation.Z = 0
	clientSession.targetLocation.X = 0
	clientSession.targetLocation.Y = 0
	clientSession.targetLocation.Z = 0
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

			clientSession.UID = res.Uid

			if res.RetCode != RETURNCODE.ReturnCode_retFail {
				clientSession.UID = res.Uid
				clientSession.RequestRoomEnterReq()
			}
		}
	case PROTOCOL.ProtocolID_RoomEnterRes:
		{
			res := &LobbyPacket.RoomEnterRes{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)

			clientSession.roomID = res.RoomID
			clientSession.roomName = res.RoomName

			log.Print(reflect.TypeOf(res))
			log.Printf("%s\n", res.String())

			if res.RetCode != RETURNCODE.ReturnCode_retFail {
				clientSession.RequestReadyForGameReq() //Considered to be loaded

				clientSession.RandomAction(1000 * time.Millisecond)
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
	case PROTOCOL.ProtocolID_RoomLeaveRes:
		{
			res := &LobbyPacket.RoomLeaveRes{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)

			log.Print(reflect.TypeOf(res))
			log.Printf("%s\n", res.String())

			if res.RetCode == RETURNCODE.ReturnCode_retOK {
				clientSession.roomID = 0
			}
		}
	case PROTOCOL.ProtocolID_FisingStartRes:
		{
			res := &LobbyPacket.FisingStartRes{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)
			log.Printf("%s\n", res.String())

			if res.RetCode == RETURNCODE.ReturnCode_retOK {
				//cancel casting.
			}
		}
	case PROTOCOL.ProtocolID_FisingStartNfy:
		{
			res := &LobbyPacket.FisingStartNfy{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)
			log.Printf("%s\n", res.String())
		}
	case PROTOCOL.ProtocolID_FisingHitNfy:
		{
			res := &LobbyPacket.FisingHitNfy{}
			err := proto.Unmarshal(buffer[:], res)
			util.ProcessError(err)
			log.Printf("%s\n", res.String())
		}
	}
}

// RequestLoginReq is ...
func (clientSession *ClientSession) RequestLoginReq(userName string) {

	req := &LobbyPacket.LoginReq{Name: userName}
	clientSession.SendPacket(PROTOCOL.ProtocolID_LoginReq, req)
}

// RequestRoomEnterReq ...
func (clientSession *ClientSession) RequestRoomEnterReq() {

	req := &LobbyPacket.RoomEnterReq{Uid: clientSession.UID}
	clientSession.SendPacket(PROTOCOL.ProtocolID_RoomEnterReq, req)
}

// RequestReadyForGameReq will request to server after client ready to play and receive packet from server(when client loaded c omplete map and game dates)
func (clientSession *ClientSession) RequestReadyForGameReq() {

	req := &LobbyPacket.ReadyForGameReq{Uid: clientSession.UID}
	clientSession.SendPacket(PROTOCOL.ProtocolID_ReadyForGameReq, req)
}

// RequestRoomLeaveReq ...
func (clientSession *ClientSession) RequestRoomLeaveReq() {

	req := &LobbyPacket.RoomLeaveReq{Uid: clientSession.UID, RoomID: clientSession.roomID}
	clientSession.SendPacket(PROTOCOL.ProtocolID_RoomLeaveReq, req)
}

// RandomMove is send moving packet to server
func (clientSession *ClientSession) RandomMove(duration time.Duration) {
	clientSession.timer = time.NewTimer(duration)
	func() {
		<-clientSession.timer.C

		go clientSession.RandomAction(duration)

		rand.NewSource(time.Now().UnixNano())

		if rand.Int31n(10) > 5 {
			clientSession.targetLocation.X += rand.Int31n(3)
			clientSession.targetLocation.Y += rand.Int31n(3)
		} else {
			clientSession.targetLocation.X -= rand.Int31n(3)
			clientSession.targetLocation.Y -= rand.Int31n(3)
		}

		//move start req
		moveStartReq := &LobbyPacket.MoveStartReq{Uid: clientSession.UID, CurrentPos: &clientSession.currentLocation, TargetPos: &clientSession.targetLocation}
		clientSession.SendPacket(PROTOCOL.ProtocolID_MoveStartReq, moveStartReq)

		clientSession.moveEndtimer = time.NewTimer(500 * time.Millisecond)
		go func() {
			<-clientSession.moveEndtimer.C

			//move end req
			moveEndReq := &LobbyPacket.MoveEndReq{Uid: clientSession.UID, TargetPos: &clientSession.targetLocation}
			clientSession.SendPacket(PROTOCOL.ProtocolID_MoveEndReq, moveEndReq)
		}()

	}()
}

// CastFishing is send fising packet to server
func (clientSession *ClientSession) CastFishing(duration time.Duration) {
	clientSession.timer = time.NewTimer(duration)
	<-clientSession.timer.C

	go clientSession.RandomAction(3000 * time.Millisecond)

	//move start req
	fishingStartReq := &LobbyPacket.FisingStartReq{AreaId: 1}
	clientSession.SendPacket(PROTOCOL.ProtocolID_FisingStartReq, fishingStartReq)

}

// RandomAction is test function package for server logic
func (clientSession *ClientSession) RandomAction(duration time.Duration) {
	rand.NewSource(time.Now().UnixNano())
	rndCase := rand.Intn(10)

	if rndCase >= 5 {
		clientSession.RandomMove(duration)
	} else {
		clientSession.CastFishing(duration)
	}
}
