package botio

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/TaKeO90/deploybot/msg"
)

const (
	sendMsgMethod string = "SendMessage"
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

// ChanRes result that we get from the channel.
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

// SendMessage structure that holds fields required for sending message.
type SendMessage struct {
	Chatid int
	Text   string
}

func sendData(data url.Values, method string) (bool, error) {
	resp, err := http.PostForm(apiurl+method, data)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}

// Send SendMessage's method for sending message returns true if everything
// ok otherwhise it returns an error.
func (s *SendMessage) Send() (bool, error) {
	val := url.Values{"chat_id": {strconv.Itoa(s.Chatid)}, "text": {s.Text}}
	isOk, err := sendData(val, sendMsgMethod)
	if err != nil {
		return false, err
	}
	if isOk {
		return true, nil
	}
	return false, nil
}
