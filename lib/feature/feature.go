package feature

import (
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type Feature struct {
}

func (f Feature) RunEvent(event slackevents.EventsAPIEvent) error {
	return nil
}

func (f Feature) RunAction(payload slack.InteractionCallback) error {
	return nil
}
