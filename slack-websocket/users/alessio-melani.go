package users

import (
	"github.com/nlopes/slack"
	"strings"
)

type MeloHandler struct{}

func (mh MeloHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) string {
	return mh.GetMessage(event.Text)
}

func (mh MeloHandler) GetMessage(receivedMessage string) string {
	if strings.Contains(receivedMessage, "erudiscimi"){

	}
	return "sei un grande"
}
