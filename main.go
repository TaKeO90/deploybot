package main

import (
	"log"

	"github.com/TaKeO90/deploybot/botio"
)

var updateid int = 0

func main() {
	c := make(chan botio.ChanRes)
	for {
		go botio.GetUpdates(c)
		result := <-c
		if result.Err != nil {
			log.Fatal(result.Err)
		}
		updateID := result.Msg.GetupdateID()
		userStuff := result.Msg.GetSenderData()
		if updateID != updateid {
			log.Println("we need to handle commands now", userStuff)
			updateid = updateID
		}
	}
}
