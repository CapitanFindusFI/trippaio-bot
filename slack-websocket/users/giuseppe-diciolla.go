package users

import (
	"github.com/nlopes/slack"
	"strings"
)

type GiuseppeHandler struct{}

func (gh GiuseppeHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) string {
	return gh.GetMessage(event.Text)
}

func (gh GiuseppeHandler) GetMessage(receivedMessage string) string {
	if strings.Contains(receivedMessage, "senza sale") {
		return "senza sale? frocio"
	}
	return "buo"
}
