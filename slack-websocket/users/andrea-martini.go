package users

import (
	"github.com/CapitanFindusFI/trippaio-bot/utils"
	"github.com/nlopes/slack"
)

type AndreaHandler struct{}

func (ah AndreaHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) []string {
	return ah.GetMessage(event.Text)
}

func (ah AndreaHandler) GetMessage(receivedMessage string) []string {
	var messages []string

	if utils.ContainsOne(receivedMessage, []string{"ciao", "buongiorno"}) {
		messages = append(messages, "ciao lillo!")
	}

	if len(messages) == 0 {
		messages = append(messages, "wow")
	}

	return messages
}
