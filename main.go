package main

import (
	"github.com/CapitanFindusFI/trippaio-bot/slack-events"
	"github.com/CapitanFindusFI/trippaio-bot/slack-websocket"
	"github.com/nlopes/slack"
	"sync"
)

var (
	api = slack.New("xoxb-7719606452-480504548341-KaEPh6Q6OwrFwnOeCeAxJIVw")
	rtm = api.NewRTM()
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go rtm.ManageConnection()
	go slack_websocket.WebsocketHandler(rtm)
	go slack_events.EventsHandler(api)
	wg.Wait()
}
