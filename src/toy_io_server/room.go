package main

import (
	"container/list"
	"log"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
	ReturnCode "packet_returncode"
	"sync"

	"github.com/golang/protobuf/proto"
)

// Room like channel or instanced dungeon. users play game in this place.
type Room struct {
	sync.Mutex
	RoomID       int32
	Name         string
	MaxUserCount int
	userList     list.List
}

func (room *Room) Init() {
	room.MaxUserCount = 128 //todo. will change setting value by data....
}

func (room *Room) Enter(session *User) bool {
	res := &LobbyPacket.RoomEnterRes{}
	res.RetCode = ReturnCode.ReturnCode_retOK

	if session.enteredRoom != nil {
		//already entered.
		res.RetCode = ReturnCode.ReturnCode_retFail
		session.SendPacket(PROTOCOL.ProtocolID_RoomEnterRes, res)
		return false
	}

	if room.addUser(session) == false {
		//full room. try enter another room
		res.RetCode = ReturnCode.ReturnCode_retFail
		session.SendPacket(PROTOCOL.ProtocolID_RoomEnterRes, res)
		return false
	}

	res.RoomID = room.RoomID
	res.RoomName = room.Name
	session.SendPacket(PROTOCOL.ProtocolID_RoomEnterRes, res)

	nfy := &LobbyPacket.RoomEnterNfy{}
	for e := room.userList.Back(); e != nil; e = e.Prev() {
		user := e.Value.(*User)

		userInfo := &LobbyPacket.UserInfo{
			Uid:      user.UID,
			Name:     user.Name,
			Location: &LobbyPacket.Location{},
		}

		nfy.UserInfoList = append(nfy.UserInfoList, userInfo)
	}

	room.Broadcast(session, PROTOCOL.ProtocolID_RoomEnterNfy, nfy)

	log.Printf("[Room] User %s Entered in %s", session.Name, room.Name)

	return true
}

func (room *Room) Leave(session *User) {
	if room.IsExistUser(session) {
		room.deleteUser(session)

		session.SetRoom(nil)

		packet := &LobbyPacket.RoomLeaveRes{}
		packet.RetCode = ReturnCode.ReturnCode_retOK
		session.SendPacket(PROTOCOL.ProtocolID_RoomLeaveRes, packet)

		nfy := &LobbyPacket.RoomLeaveNfy{}
		for e := room.userList.Back(); e != nil; e = e.Prev() {
			user := e.Value.(*User)

			userInfo := &LobbyPacket.UserInfo{
				Uid:      user.UID,
				Name:     user.Name,
				Location: &LobbyPacket.Location{},
			}

			nfy.UserInfoList = append(nfy.UserInfoList, userInfo)
		}

		room.Broadcast(session, PROTOCOL.ProtocolID_RoomLeaveNfy, nfy)

		log.Printf("[Room] User %s left from %s", session.Name, room.Name)
	}
}

func (room *Room) SendUserList(session *User) {
}

func (room *Room) Broadcast(sender *User, protocolID PROTOCOL.ProtocolID, pb proto.Message) {
	if room == nil {
		return //it occur test case only.
	}

	for e := room.userList.Back(); e != nil; e = e.Prev() {
		user := e.Value.(*User)
		if user != nil && user != sender {
			user.SendPacket(protocolID, pb)
		}
	}
}

func (room *Room) addUser(session *User) bool {
	if room.userList.Len() > room.MaxUserCount {
		return false
	}

	room.Lock()
	room.userList.PushBack(session)
	room.Unlock()

	return true
}

func (room *Room) deleteUser(session *User) bool {

	room.Lock()
	for e := room.userList.Back(); e != nil; e = e.Prev() {
		if e.Value == session {
			room.userList.Remove(e)
			room.Unlock()
			return true
		}
	}
	room.Unlock()
	return false
}

func (room *Room) IsExistUser(Session *User) bool {
	room.Lock()
	for e := room.userList.Back(); e != nil; e = e.Prev() {
		if e.Value == Session {
			room.Unlock()
			return true
		}
	}

	room.Unlock()
	return false
}

func (room *Room) GetRoomNumber() int32 {
	return room.RoomID
}

func (room *Room) SetMaxUserCount(userCount int) {
	room.MaxUserCount = userCount
}

func (room *Room) GetMaxUserCount() int {

	return room.MaxUserCount
}

func (room *Room) SetRoomName(name string) bool {
	return room.SetRoomName(name)
}

func (room *Room) GetRoomName(Session *User) string {
	return room.Name
}

func (room *Room) IsFullRoom() bool {
	return room.userList.Len() > room.MaxUserCount
}
