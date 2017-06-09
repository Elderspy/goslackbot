package helpers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync/atomic"
)

type responseRtmStart struct {
	OK    bool         `json:"ok"`
	Error string       `json:"error"`
	URL   string       `json:"url"`
	Self  responseSelf `json:"self"`
}

type Message struct {
	ID      uint64 `json:"id"`
	Type    string `json:"type"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type responseSelf struct {
	ID string `json:"id"`
}

func SlackConnect(token string) (*websocket.Conn, string) {
	url := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", token)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		panic(fmt.Errorf("Slack did not return 200 response was %s", strconv.Itoa(response.StatusCode)))
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		panic(err)
	}

	var respObject responseRtmStart
	err = json.Unmarshal(body, &respObject)
	if err != nil {
		panic(err)
	}
	if !respObject.OK {
		err = fmt.Errorf("Slack error :%s", respObject.Error)
		panic(err)
	}

	wsurl := respObject.URL
	id := respObject.Self.ID

	ws, err := websocket.Dial(wsurl, "", "https://api.slack.com/")
	if err != nil {
		panic(err)
	}
	return ws, id
}

func GetMessage(ws *websocket.Conn) (m Message, err error) {
	err = websocket.JSON.Receive(ws, &m)
	return
}

var counter uint64

func SendMessage(ws *websocket.Conn, content Message) error {
	content.ID = atomic.AddUint64(&counter, 1)
	return websocket.JSON.Send(ws, content)
}
