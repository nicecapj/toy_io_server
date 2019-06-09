package main

import (
	"container/list"
	"sync"
)

type RoomManager struct {
	sync.Mutex
	roomList *list.List
}

var roomManagerInstance *RoomManager
var roomManagerOnce sync.Once

func GetRoomManager() *RoomManager {
	roomManagerOnce.Do(func() {
		roomManagerInstance = &RoomManager{}
		roomManagerInstance.Init()
	})

	return roomManagerInstance
}

// Init ...
func (roomManager *RoomManager) Init() {
	roomManager.roomList = list.New()
}

func (roomManager *RoomManager) GetLeisuerlyRoom() *Room {
	for e := roomManager.roomList.Front(); e != nil; e = e.Next() {
		//var room Room = e.Value

		// if room.IsFullRoom() == false {
		// 	return &room
		// }
	}

	return roomManager.AddRoom()
}

func (roomManager *RoomManager) AddRoom() *Room {
	room := new(Room)
	roomManager.roomList.PushBack(room)

	return room
}
