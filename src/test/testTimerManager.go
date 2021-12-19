package main

import (
	"log"
	"network"
	"time"
)

//-----------------------------------------------------------------------------
func main() {
	//timer
	timerManager := network.GetTimerManager()
	timerManager.AddTimer(3000, func() {
		log.Println("call timer")
	})
	log.Println("call timer end")

	timeTemp := time.NewTimer(time.Millisecond * 10000)
	<-timeTemp.C //블록하지 말고, 비동기로 처리
}
