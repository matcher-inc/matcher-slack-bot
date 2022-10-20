package mSlack

import (
	"bytes"
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
		FeaturePath: slash.Command[1:],
		ChannelID:   slash.ChannelID,
		UserID:      slash.UserID,
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
				FeaturePath: strings.Split(data.Text, " ")[1],
				UserID:      data.User,
				ChannelID:   data.Channel,
			}
			eventType = AppMentionEvent
			return
		}
	}
	err = errors.New("Failed Parse Event")
	return
}

func ParseAction(r *http.Request) (params *RequestParams, err error) {
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

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var payload *slack.InteractionCallback
	err = json.Unmarshal([]byte(r.FormValue("payload")), &payload)
	if err != nil {
		return
	}

	switch payload.Type {
	case slack.InteractionTypeBlockActions:
		if len(payload.ActionCallback.BlockActions) == 0 {
			err = errors.New("Invalid action")
			return
		}
		action := payload.ActionCallback.BlockActions[0]
		path := strings.Split(action.BlockID, ":")
		params = &RequestParams{
			FeaturePath: path[0],
			ActionPath:  path[1],
			UserID:      payload.User.ID,
			ChannelID:   payload.Channel.ID,
		}
		return
	case slack.InteractionTypeViewSubmission:
		path := strings.Split(payload.View.CallbackID, ":")
		params = &RequestParams{
			FeaturePath: path[0],
			ActionPath:  path[1],
			UserID:      payload.User.ID,
			ChannelID:   payload.Channel.ID,
		}
		return
	}
	err = errors.New("Failed Parse Action")
	return
}
