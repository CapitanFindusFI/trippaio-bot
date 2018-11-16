package users

import (
	"github.com/nlopes/slack"
)

type BiniHandler struct{}

func (bh BiniHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) string {
	return bh.GetMessage(event.Text)
}

func (bh BiniHandler) GetMessage(receivedMessage string) string {
	return "*ALBEEEERTOOOOO*"
}
