package msg

import (
	"encoding/json"
)

// Message structure that represent a received message from telegram.
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
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

// ParseJson unmarshal json that we get from webhook
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

// SenderData
type SenderData struct {
	Id        int
	IsBot     bool
	Firstname string
	Lastname  string
	Username  string
}

// TextMsgData
type TextMsgData struct {
	Date     int
	Text     string
	Entities []Entity
}

// Entity
type Entity struct {
	Type   string
	Offset int
	Length int
}

// GetSenderData get the sender informations from Message.
func (m Message) GetSenderData() SenderData {
	id := m.Msg.From.Id
	isBot := m.Msg.From.IsBot
	firstname := m.Msg.From.Firstname
	lastname := m.Msg.From.Lastname
	username := m.Msg.From.Username
	return *(&SenderData{id, isBot, firstname, lastname, username})
}

// GetTxtMsgData get text message information from Message.
func (m Message) GetTxtMsgData() TextMsgData {
	var entities []Entity
	date, text := m.Msg.Date, m.Msg.Text
	for _, e := range m.Msg.Entities {
		ent := &Entity{e.Type, e.Offset, e.Length}
		entities = append(entities, *ent)
	}
	return *(&TextMsgData{date, text, entities})
}
