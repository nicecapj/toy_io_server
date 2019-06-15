package main

import (
	"container/list"
	LobbyPacket "packet_lobby"
	PROTOCOL "packet_protocol"
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
	if room.addUser(session) == false {
		return false
	}

	packet := &LobbyPacket.RoomEnterNfy{}

	for e := room.userList.Back(); e != nil; e = e.Prev() {
		user := e.Value.(*User)

		userInfo := &LobbyPacket.UserInfo{
			Uid:      user.Uid,
			Name:     user.Name,
			Location: &LobbyPacket.Location{},
		}

		packet.UserInfoList = append(packet.UserInfoList, userInfo)
	}

	room.Broadcast(session, PROTOCOL.ProtocolID_RoomEnterNfy, packet)

	return true
}

func (room *Room) Leave(session *User) {
	room.deleteUser(session)
}

func (room *Room) SendUserList(session *User) {
}

func (room *Room) Broadcast(sender *User, protocolID PROTOCOL.ProtocolID, pb proto.Message) {
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
	for e := room.userList.Back(); e != nil; e = e.Prev() {
		if e.Value == Session {
			return true
		}
	}

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
