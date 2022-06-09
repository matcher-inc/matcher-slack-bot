package listner

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

type ActionError error

var (
	ActionStandardError ActionError = errors.New("error 1")
)

type ActionRoute interface {
	match(payload slack.InteractionCallback) bool
	handle(payload slack.InteractionCallback) ActionError
}

type InteractionTypeBlockActionRoute struct {
	ActionId string
	Handler  func(payload slack.InteractionCallback) ActionError
}

func (a InteractionTypeBlockActionRoute) match(payload slack.InteractionCallback) bool {
	switch payload.Type {
	case slack.InteractionTypeBlockActions:
		if len(payload.ActionCallback.BlockActions) == 0 {
			return false
		}
		action := payload.ActionCallback.BlockActions[0]
		return action.BlockID == a.ActionId
	}
	return false
}

func (a InteractionTypeBlockActionRoute) handle(payload slack.InteractionCallback) (error ActionError) {
	return a.Handler(payload)
}

func ListenAction(routes []ActionRoute) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		verifier, err := slack.NewSecretsVerifier(r.Header, os.Getenv("SLACK_SIGNING_SECRET"))
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
		for _, route := range routes {
			if route.match(*payload) {
				error := route.handle(*payload)
				if error != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
		}
	}
}
