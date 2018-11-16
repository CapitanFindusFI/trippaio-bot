package slack_events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
	"net/http"
)

func EventsHandler(api *slack.Client) {
	http.HandleFunc("/events-endpoint", func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()
		eventsAPIEvent, e := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionVerifyToken(&slackevents.TokenComparator{VerificationToken: "TOKEN"}))
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if eventsAPIEvent.Type == slackevents.URLVerification {
			var r *slackevents.ChallengeResponse
			err := json.Unmarshal([]byte(body), &r)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "text")
			w.Write([]byte(r.Challenge))
		}
		if eventsAPIEvent.Type == slackevents.CallbackEvent {
			innerEvent := eventsAPIEvent.InnerEvent
			switch ev := innerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				api.PostMessage(ev.Channel, slack.MsgOptionText("Yes, hello.", false))
			}
		}
	})
	fmt.Println("[INFO] Server listening")
	http.ListenAndServe(":8000", nil)
}