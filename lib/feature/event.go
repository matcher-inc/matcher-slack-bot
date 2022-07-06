package feature

import (
	"errors"

	"github.com/slack-go/slack/slackevents"
)

type EventType string

const (
	AppMentionEvent EventType = "AppMentionEvent"
)

type Event interface {
	match(event slackevents.EventsAPIEvent) bool
	handle(event slackevents.EventsAPIEvent) error
}

type MentionEvent struct {
	Callback func(*slackevents.AppMentionEvent) error
}

func (e MentionEvent) match(event slackevents.EventsAPIEvent) bool {
	switch event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		return true
	}
	return false
}

func (e MentionEvent) handle(event slackevents.EventsAPIEvent) error {
	switch data := event.InnerEvent.Data.(type) {
	case *slackevents.AppMentionEvent:
		return e.Callback(data)
	}
	return errors.New("エラー")
}

func (f Feature) RunEvent(event slackevents.EventsAPIEvent) error {
	if f.Event.match(event) {
		return f.Event.handle(event)
	}
	return errors.New("タイプが一致しません。")
}
