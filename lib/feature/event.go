package feature

import "github.com/slack-go/slack/slackevents"

type Event struct {
	Callback func(slackevents.EventsAPIEvent) error
}

func (f Feature) RunEvent(event slackevents.EventsAPIEvent) error {
	return f.Event.Callback(event)
}
