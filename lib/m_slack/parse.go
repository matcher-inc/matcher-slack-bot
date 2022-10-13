package mSlack

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func ParseSlash(r *http.Request) (params *RequestParams, err error) {
	slash, err := slack.SlashCommandParse(r)
	if err != nil {
		return
	}

	err = VerificateSlashToken(slash.Token)
	if err != nil {
		return
	}

	params = &RequestParams{
		RequestKey: slash.Command[1:],
		ChannelID:  slash.ChannelID,
		UserID:     slash.UserID,
	}
	return
}

func ParseEvent(r *http.Request) (params *RequestParams, requestBody []byte, eventType EventType, err error) {
	verifier, err := verificateSigningSecret(r)
	if err != nil {
		return
	}

	bodyReader := io.TeeReader(r.Body, &verifier)
	body, err := ioutil.ReadAll(bodyReader)
	if err != nil {
		return
	}

	err = verifier.Ensure()
	if err != nil {
		return
	}

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		return
	}

	switch eventsAPIEvent.Type {
	case slackevents.URLVerification:
		params = &RequestParams{}
		requestBody = body
		eventType = URLVerification
		return
	case slackevents.CallbackEvent:
		switch data := eventsAPIEvent.InnerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			params = &RequestParams{
				RequestKey: strings.Split(data.Text, " ")[1],
				UserID:     data.User,
				ChannelID:  data.Channel,
			}
			eventType = AppMentionEvent
			return
		}
	}
	err = errors.New("Failed Parse Event")
	return
}
