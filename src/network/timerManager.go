package network

import (
	"sync"
	"time"
)

// RoomManager ...
type TimerManager struct {
	sync.Mutex
}

var timerManagerInstance *TimerManager
var timerManagerOnce sync.Once

// GetRoomManager is singleton
func GetTimerManager() *TimerManager {
	timerManagerOnce.Do(func() {
		timerManagerInstance = &TimerManager{}
		timerManagerInstance.Init()
	})

	return timerManagerInstance
}

// Init ...
func (manager *TimerManager) Init() {

}

func (manager *TimerManager) AddTimer(msec time.Duration, callback func()) {
	timeTemp := time.NewTimer(time.Millisecond * msec)

	//<-timeTemp.C	//block
	go func() {
		<-timeTemp.C //async
		callback()
	}()
}
