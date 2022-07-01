package web

import (
	"bytes"
	"encoding/json"
	"go-bot-test/config/routes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

func handleAction(w http.ResponseWriter, r *http.Request) {
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
	for _, route := range routes.Rounting {
		if actionIsMatchingToRoute(*payload, route) {
			error := route.Feature.RunAction(*payload)
			if error != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}

func actionIsMatchingToRoute(payload slack.InteractionCallback, route routes.Route) bool {
	switch payload.Type {
	case slack.InteractionTypeBlockActions:
		if len(payload.ActionCallback.BlockActions) == 0 {
			return false
		}
		action := payload.ActionCallback.BlockActions[0]
		path := strings.Split(action.BlockID, ":")[0]
		return path == route.Path
	}
	return false
}
