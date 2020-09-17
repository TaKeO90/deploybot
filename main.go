package main

import (
	"log"

	"github.com/TaKeO90/deploybot/botio"
	"github.com/TaKeO90/deploybot/cmdhandler"
)

var updateid int = 0

func main() {
	c := make(chan botio.ChanRes)
	for {
		//get updates from the webhook
		go botio.GetUpdates(c)
		result := <-c
		if result.Err != nil {
			log.Fatal(result.Err)
		}
		//parse data that we get from the webhook
		updateID := result.Msg.GetupdateID()
		userStuff := result.Msg.GetSenderData()
		msgStuff := result.Msg.GetTxtMsgData()
		if updateID != updateid {
			//Need to give the cmdhandler here sender stuff & text msg from webhook
			var (
				text    string
				fstname string
				lstname string
				id      int
			)
			if msgStuff["text"] != nil || userStuff["firstname"] != nil || userStuff["lastname"] != nil || userStuff["Id"] != nil {
				text, fstname, lstname, id = msgStuff["text"].(string), userStuff["firstname"].(string), userStuff["lastname"].(string), userStuff["Id"].(int)
			}
			cmd := cmdhandler.NewCmd(text, fstname, lstname, id)
			cmd.HandleCmd()
			updateid = updateID
		}
	}
}
