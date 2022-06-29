package listner

import (
	"encoding/json"
	"errors"
	"go-bot-test/config/env"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type EventError error

var (
	EventStandardError EventError = errors.New("error 1")
)

type EventRoute interface {
	match(event slackevents.EventsAPIEvent) bool
	handle(event slackevents.EventsAPIEvent) EventError
}

type MentionEventRoute struct {
	Command string
	Handler func(event slackevents.AppMentionEvent) EventError
}

func (h MentionEventRoute) match(event slackevents.EventsAPIEvent) bool {
	switch event.Type {
	case slackevents.CallbackEvent:
		switch data := event.InnerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			message := strings.Split(data.Text, " ")
			command := message[1]
			return h.Command == command
		}
	}
	return false
}

func (h MentionEventRoute) handle(event slackevents.EventsAPIEvent) (error EventError) {
	switch event.Type {
	case slackevents.CallbackEvent:
		switch data := event.InnerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			return h.Handler(*data)
		}
	}
	return
}

func verificateUrl(w http.ResponseWriter, body []byte) {
	var res *slackevents.ChallengeResponse
	if err := json.Unmarshal(body, &res); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte(res.Challenge)); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ListenEvent(routes []EventRoute) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		verifier, err := slack.NewSecretsVerifier(r.Header, env.SLACK_SIGNING_SECRET)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bodyReader := io.TeeReader(r.Body, &verifier)
		body, err := ioutil.ReadAll(bodyReader)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := verifier.Ensure(); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		switch eventsAPIEvent.Type {
		case slackevents.URLVerification:
			verificateUrl(w, body)
		case slackevents.CallbackEvent:
			for _, route := range routes {
				if route.match(eventsAPIEvent) {
					error := route.handle(eventsAPIEvent)
					if error != nil {
						w.WriteHeader(http.StatusInternalServerError)
						return
					}
				}
			}
		}
	}
}
