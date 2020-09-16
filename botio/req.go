package botio

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/TaKeO90/deploybot/msg"
)

var (
	apitoken   = os.Getenv("BOT_TOKEN")
	webhookurl = os.Getenv("BOT_WEBHOOK_URL")
	apiurl     = os.Getenv("BOT_API_URL")
)

func getWebhookUpdates() ([]byte, int, error) {
	resp, err := http.Get(webhookurl)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return resBody, resp.StatusCode, nil
}

type ChanRes struct {
	Msg msg.Message
	Err error
}

// GetUpdates dial to the webhook url and get data.
// and parse it using msg package
func GetUpdates(c chan ChanRes) {
	result := new(ChanRes)
	respR, statusCode, err := getWebhookUpdates()
	if err != nil {
		result.Err = err
		c <- *result
	}
	if statusCode == 200 {
		m, err := msg.ParseJson(respR)
		if err != nil {
			result.Err = err
			c <- *result
		}
		result.Msg, result.Err = m, nil
		c <- *result
	}
}
