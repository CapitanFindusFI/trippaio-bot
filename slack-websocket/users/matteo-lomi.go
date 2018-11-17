package users

import "github.com/nlopes/slack"

type LomiHandler struct{}

func (lh LomiHandler) Handle(rtm *slack.RTM, event *slack.MessageEvent) []string {
	return lh.GetMessage(event.Text)
}

func (lh LomiHandler) GetMessage(receivedMessage string) []string {
	var messages []string

	if len(messages) == 0 {
		messages = append(messages, "_Lomi_ sei enorme")
	}

	return messages
}
