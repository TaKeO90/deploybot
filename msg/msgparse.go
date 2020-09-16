package msg

import (
	"encoding/json"
)

type Message struct {
	UpdateId int `json:"update_id"`
	Msg      struct {
		MessageID int `json:"message_id"`
		From      struct {
			Id        int    `json:"id"`
			IsBot     bool   `json:"is_bot"`
			Firstname string `json:"first_name"`
			Lastname  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"from"`
		ChatC struct {
			Id        int    `json:"id"`
			Firstname string `json:"first_name"`
			Lastname  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

func ParseJson(data []byte) (Message, error) {
	msg := new(Message)
	err := json.Unmarshal(data, msg)
	if err != nil {
		return *msg, err
	}
	return *msg, nil
}

// GetupdateID return the update it from Message.
func (m Message) GetupdateID() int {
	return m.UpdateId
}

// GetSenderData get the sender informations from Message.
func (m Message) GetSenderData() map[string]interface{} {
	resM := map[string]interface{}{
		"Id":        m.Msg.From.Id,
		"IsBot":     m.Msg.From.IsBot,
		"firstname": m.Msg.From.Firstname,
		"lastname":  m.Msg.From.Lastname,
		"username":  m.Msg.From.Username,
	}
	return resM
}

// GetTxtMsgData get text message information from Message.
func (m Message) GetTxtMsgData() map[string]interface{} {
	resT := map[string]interface{}{
		"date":     m.Msg.Date,
		"text":     m.Msg.Text,
		"entities": m.Msg.Entities,
	}
	return resT
}
