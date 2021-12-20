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

	log.Println("call timer")
	timerManager.AddTimer(3000, func() {
		log.Println("called function after 3sec")
	})

	log.Println("wait 10sec")
	timeTemp := time.NewTimer(time.Millisecond * 10000)
	<-timeTemp.C //wait for finish
}
