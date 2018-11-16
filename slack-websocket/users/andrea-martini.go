package users

import "github.com/nlopes/slack"

type AndreaHandler struct{}

func (ah AndreaHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) string {
	return ah.GetMessage(event.Text)
}

func (ah AndreaHandler) GetMessage(receivedMessage string) string {
	return "wow"
}
