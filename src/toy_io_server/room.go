package main

import "container/list"

type Room struct {
	Name         string
	UserList     list.List
	MaxUserCount int32
}

func (room *Room) Enter(Session *User) bool {
	return true
}

func (room *Room) Leave(Session *User) {

}

func (room *Room) SendUserList(Session *User) {
}

func (room *Room) Broadcast(Session *User) {
}

func (room *Room) AddUser(Session *User) bool {
	return true
}

func (room *Room) DeleteUser(Session *User) {

}

func (room *Room) IsExistUser(Session *User) bool {
	return true
}

func (room *Room) GetRoomNumber() int32 {
	return 0
}

func (room *Room) SetMaxUserCount(userCount int32) {
	room.MaxUserCount = userCount
}

func (room *Room) GetMaxUserCount() int32 {
	return 20
}

func (room *Room) SetRoomName(name string) bool {
	return true
}

func (room *Room) GetRoomName(Session *User) string {
	return room.Name
}

func (room *Room) IsFullRoom() bool {
	return true
}
