package slack_websocket

import "github.com/nlopes/slack"

type UserHandler interface {
	Handle(rtm *slack.RTM, event *slack.MessageEvent) []string
	GetMessage(receivedText string) []string
}
