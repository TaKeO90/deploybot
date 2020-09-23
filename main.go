package main

import (
	"log"

	"github.com/TaKeO90/deploybot/botio"
	"github.com/TaKeO90/deploybot/cmdhandler"
	"github.com/TaKeO90/deploybot/msg"
)

//TODO: log package logs info error into the bot log file.

var updateid int = 0

func main() {
	c := make(chan botio.ChanRes)
	for {
		result, err := getWebhookUpdates(c)
		if err != nil {
			log.Fatal(err)
		}
		//parse data that we get from the webhook
		cData := getDataFromResult(result)
		if cData.updateID != updateid {
			cData.handleCmds()
			updateid = cData.updateID
		}
	}
}

type clientData struct {
	updateID    int
	senderData  map[string]interface{}
	textMsgData map[string]interface{}
}

func typeAssertClienData(t, fstN, lstN *string, id *int, text, firstname, lastname, Id interface{}) {
	if firstname != nil || lastname != nil || Id != nil || text != nil {
		*t, *fstN, *lstN, *id = text.(string), firstname.(string), lastname.(string), Id.(int)
	}
}

func (c clientData) handleCmds() {
	var (
		text    string
		fstname string
		lstname string
		id      int
	)
	typeAssertClienData(&text, &fstname, &lstname, &id, c.textMsgData["text"], c.senderData["firstname"], c.senderData["lastname"], c.senderData["Id"])
	cmd := cmdhandler.NewCmd(text, fstname, lstname, id)
	cmd.HandleCmd()
}

func getWebhookUpdates(c chan botio.ChanRes) (msg.Message, error) {
	go botio.GetUpdates(c)
	result := <-c
	if result.Err != nil {
		return result.Msg, result.Err
	}
	return result.Msg, nil
}

func getDataFromResult(message msg.Message) clientData {
	cData := new(clientData)
	cData.updateID = message.GetupdateID()
	cData.senderData = message.GetSenderData()
	cData.textMsgData = message.GetTxtMsgData()
	return *cData
}
