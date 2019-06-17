package main

import (
	"container/list"
	"strconv"
	"sync"
)

// RoomManager ...
type RoomManager struct {
	sync.Mutex
	roomList         *list.List
	currentRoomIndex int32
}

var roomManagerInstance *RoomManager
var roomManagerOnce sync.Once

// GetRoomManager is singleton
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

// GetLeisuerlyRoom return a room that is not full
func (roomManager *RoomManager) GetLeisuerlyRoom() *Room {
	roomManager.Lock()
	for e := roomManager.roomList.Front(); e != nil; e = e.Next() {

		room := e.Value.(*Room)

		if room.IsFullRoom() == false {
			roomManager.Unlock()
			return room
		}
	}
	roomManager.Unlock()

	room := roomManager.CreateRoom()

	return room
}

// FindRoom return a finded room
func (roomManager *RoomManager) FindRoom(roomID int32) *Room {
	roomManager.Lock()
	for e := roomManager.roomList.Front(); e != nil; e = e.Next() {

		room := e.Value.(*Room)

		if room.RoomID == roomID {
			roomManager.Unlock()
			return room
		}
	}
	roomManager.Unlock()

	return nil
}

// CreateRoom ...
func (roomManager *RoomManager) CreateRoom() *Room {
	roomManager.Lock()

	room := new(Room)
	room.Init()
	roomManager.roomList.PushBack(room)

	roomManager.currentRoomIndex = roomManager.currentRoomIndex + 1
	room.RoomID = roomManager.currentRoomIndex
	room.Name = "Room" + strconv.Itoa(int(room.RoomID))

	roomManager.Unlock()

	return room
}
