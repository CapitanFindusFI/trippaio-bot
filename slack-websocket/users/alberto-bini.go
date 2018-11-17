package users

import (
	"github.com/CapitanFindusFI/trippaio-bot/utils"
	"github.com/nlopes/slack"
)

type BiniHandler struct{}

func (bh BiniHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) []string {
	return bh.GetMessage(event.Text)
}

func (bh BiniHandler) GetMessage(receivedMessage string) []string {
	var messages []string

	if utils.ContainsOne(receivedMessage, []string{"ciao", "buongiorno"}) {
		messages = append(messages, "ciao lillo!")
	}

	if len(messages) == 0 {
		messages = append(messages, "*Albeeeeeeerto*")
	}

	return messages
}
