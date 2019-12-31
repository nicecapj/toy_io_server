package main

import (
	"math/rand"
	"sync"
	"time"
)

// RoomManager ...
type FishingManager struct {
	sync.Mutex
}

var fishingManagerInstance *FishingManager
var fishingManagerOnce sync.Once

// GetRoomManager is singleton
func GetFishingManager() *FishingManager {
	fishingManagerOnce.Do(func() {
		fishingManagerInstance = &FishingManager{}
		fishingManagerInstance.Init()
	})

	return fishingManagerInstance
}

// Init ...
func (manager *FishingManager) Init() {

}

// CatchFish
// 지역 타입은 user에 저장되 있어야 하지만, 패스!
// return 0 : fish id( if 0 failed to catch)
// return 1 : CM
func (manager *FishingManager) CatchFish(areaType int32) (int32, int32) {
	//manager.Lock()

	succeed := false
	var spawnPackId int32
	switch areaType {
	case 1:
		{
			succeed = manager.Win(90 * 1000)
			spawnPackId = 1
		}
	case 2:
		{
			succeed = manager.Win(80 * 1000)
			spawnPackId = 2
		}
	case 3:
		{
			succeed = manager.Win(70 * 1000)
			spawnPackId = 3
		}
	}
	//manager.Unlock()

	if succeed {
		manager.SelectFish(spawnPackId)
	}

	return 3, 3
}

// 1/1000 = 1% = 1000
func (manager *FishingManager) Win(prob int32) bool {

	rand.NewSource(time.Now().UnixNano())

	decideValue := rand.Int31n(100000)

	if decideValue <= prob {
		return true
	}

	return false
}

func (manager *FishingManager) SelectFish(spawnPackId int32) int32 {
	//데이터라 치고
	rand.NewSource(time.Now().UnixNano())
	switch spawnPackId {
	case 1:
		{
			return rand.Int31n(10)
		}
	case 2:
		{
			return rand.Int31n(20)
		}
	case 3:
		{
			return rand.Int31n(30)
		}
	}

	return 99999999
}
