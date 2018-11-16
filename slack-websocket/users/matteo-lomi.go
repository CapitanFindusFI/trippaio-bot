package users

import "github.com/nlopes/slack"

type LomiHandler struct{}

func (lh LomiHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) string {
	return lh.GetMessage(event.Text)
}

func (lh LomiHandler) GetMessage(receivedMessage string) string {
	return "_Lomi_ sei un cazzo di grande"
}

