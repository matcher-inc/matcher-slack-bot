package web

import (
	"encoding/json"
	"errors"
	"go-bot-test/config/env"
	"go-bot-test/config/routes"
	"go-bot-test/lib/feature"
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
		params, err := parseEventRequest(eventsAPIEvent)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		route, err := routes.GetRoute(params.RequestKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		route.Feature.RunEvent(route.Path, *params)
	}
}

func parseEventRequest(event slackevents.EventsAPIEvent) (params *feature.EventParams, err error) {
	switch event.Type {
	case slackevents.CallbackEvent:
		switch data := event.InnerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			params = &feature.EventParams{
				Type:       feature.AppMentionEvent,
				RequestKey: strings.Split(data.Text, " ")[1],
				UserID:     data.User,
				ChannelID:  data.Channel,
			}
			return
		}
	}
	err = errors.New("タイプが一致しません。")
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
