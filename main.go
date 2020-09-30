package main

import (
	"log"

	"github.com/TaKeO90/deploybot/botio"
	"github.com/TaKeO90/deploybot/cmdhandler"
	"github.com/TaKeO90/deploybot/logger"
	"github.com/TaKeO90/deploybot/msg"
)

// MAIN PACKAGE ARTCHITECHTURE
// - GET UPDATES FROM WEBHOOOK. [X]
// - PARSE THOSE UPDATES IN A JSON FORMAT. [X]
// - GIVE THE PARSED UPDATES TO THE COMMAND HANDLER. -> COMMAND HANDLER ANSWERS THE USER USING BOTIO PACKAGE. [X]
// - GIVE THE PARSED UPDATES TO THE LOGGER. [X]
// - NEED SERVICE (DEPLOY SERVICE & GET STATS OF THE DEPLOYMENT)[ ]

var updateid int = 0

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

type clientData struct {
	updateID    int
	senderData  msg.SenderData
	textMsgData msg.TextMsgData
}

func (c clientData) handleCmds() error {
	cmd := cmdhandler.NewCmd(c.textMsgData.Text, c.senderData)
	err := cmd.HandleCmd()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	c := make(chan botio.ChanRes)
	writer, err := logger.OpenLogFile()
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()
	for {
		result, err := getWebhookUpdates(c)
		if err != nil {
			errLogger := logger.NewErrorLog(writer, err.Error())
			errLogger.ErrLog()
		}
		//parse data that we get from the webhook
		cData := getDataFromResult(result)
		msgLogger := logger.NewMsgLog(writer, cData.senderData,
			cData.textMsgData)
		if cData.updateID != updateid {
			err := cData.handleCmds()
			msgLogger.MsgLog()
			if err != nil {
				errLogger := logger.NewErrorLog(writer, err.Error())
				errLogger.ErrLog()
			}
			updateid = cData.updateID
		}
	}
}
