package users

import (
	"github.com/CapitanFindusFI/trippaio-bot/utils"
	"github.com/nlopes/slack"
)

type MeloHandler struct{}

func (mh MeloHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) []string {
	return mh.GetMessage(event.Text)
}

func (mh MeloHandler) GetMessage(receivedMessage string) []string {
	var messages []string

	if utils.ContainsOne(receivedMessage, []string{"ciao", "buongiorno"}) {
		messages = append(messages, "ciao ragazzi!")
	}

	if len(messages) == 0 {
		messages = append(messages, "sei veramente un grande")
	}

	return messages
}
