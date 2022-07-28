package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-bot-test/config/env"
	"go-bot-test/config/routes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/slack-go/slack"
)

func handleAction(w http.ResponseWriter, r *http.Request) {
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

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	var payload *slack.InteractionCallback
	if err := json.Unmarshal([]byte(r.FormValue("payload")), &payload); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	path, err := parseAction(*payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	route, err := routes.GetRoute(*path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	route.Feature.RunAction(*path, *payload)
}

func parseAction(payload slack.InteractionCallback) (*string, error) {
	switch payload.Type {
	case slack.InteractionTypeBlockActions:
		if len(payload.ActionCallback.BlockActions) == 0 {
			return nil, errors.New("Invalid action")
		}
		action := payload.ActionCallback.BlockActions[0]
		path := strings.Split(action.BlockID, ":")[0]
		return &path, nil
	case slack.InteractionTypeViewSubmission:
		path := strings.Split(payload.View.CallbackID, ":")[0]
		return &path, nil
	}
	return nil, errors.New("Invalid action")
}
