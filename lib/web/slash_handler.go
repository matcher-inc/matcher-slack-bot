package web

import (
	"go-bot-test/config/env"
	"log"
	"net/http"

	"github.com/slack-go/slack"
)

func handleSlash(w http.ResponseWriter, r *http.Request) {
	body, err := slack.SlashCommandParse(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if body.Token != env.SLACK_VERIFICATION_TOKEN {
		log.Println("Failed Verification Token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
