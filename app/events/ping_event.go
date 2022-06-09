package events

import (
	"go-bot-test/lib/api"
	"go-bot-test/lib/listner"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func PingEvent(event slackevents.AppMentionEvent) (error listner.EventError) {
	if _, _, err := api.Shared.PostMessage(event.Channel, slack.MsgOptionText("pong", false)); err != nil {
		log.Println(err)
		return listner.EventStandardError
	}
	return
}
