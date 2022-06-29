package web

import (
	"encoding/json"
	"go-bot-test/config/env"
	"go-bot-test/config/routes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func handleEvent(w http.ResponseWriter, r *http.Request) {
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
		for _, route := range routes.Rounting {
			if eventIsMatchingToRoute(eventsAPIEvent, route) {
				error := route.Feature.RunEvent(eventsAPIEvent)
				if error != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
	}
}

func eventIsMatchingToRoute(event slackevents.EventsAPIEvent, route routes.Route) bool {
	switch event.Type {
	case slackevents.CallbackEvent:
		switch data := event.InnerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			message := strings.Split(data.Text, " ")
			command := message[1]
			return route.Path == command
		}
	}
	return false
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
