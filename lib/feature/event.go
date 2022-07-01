package feature

import "github.com/slack-go/slack/slackevents"

type Event struct {
	Callback func(slackevents.EventsAPIEvent) (*View, error)
}

func (f Feature) RunEvent(event slackevents.EventsAPIEvent) (*View, error) {
	return f.Event.Callback(event)
}
