package users

import (
	"github.com/nlopes/slack"
	"strings"
)

type GiuseppeHandler struct{}

func (gh GiuseppeHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) []string {
	return gh.GetMessage(event.Text)
}

func (gh GiuseppeHandler) GetMessage(receivedMessage string) []string {
	var messages []string
	if strings.Contains(receivedMessage, "bnbn") {
		messages = append(messages, "*ROOTER*")
	}
	if strings.Contains(receivedMessage, "senza sale") {
		messages = append(messages, "senza sale? o che sei diventao buo?")
	}

	if len(messages) == 0 {
		messages = append(messages, "buo")
	}

	return messages
}
