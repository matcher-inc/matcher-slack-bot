package web

import (
	"go-bot-test/config/env"
	"go-bot-test/config/routes"
	"go-bot-test/lib/feature"
	"log"
	"net/http"

	"github.com/slack-go/slack"
)

func handleSlash(w http.ResponseWriter, r *http.Request) {
	slash, err := slack.SlashCommandParse(r)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if slash.Token != env.SLACK_VERIFICATION_TOKEN {
		log.Println("Failed Verification Token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := parseSlashEvent(slash)
	route, err := routes.GetRoute(params.RequestKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	route.Feature.RunEvent(route.Path, *params)
}

func parseSlashEvent(slash slack.SlashCommand) *feature.EventParams {
	return &feature.EventParams{
		Type:       feature.SlashEvent,
		RequestKey: slash.Command[1:],
		ChannelID:  slash.ChannelID,
		UserID:     slash.UserID,
	}
}
