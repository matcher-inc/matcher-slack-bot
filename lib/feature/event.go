package feature

import (
	"errors"
)

type EventType string

const (
	AppMentionEvent EventType = "AppMentionEvent"
	SlashEvent      EventType = "SlashEvent"
)

type EventParams struct {
	Type       EventType
	RequestKey string
	UserID     string
	ChannelID  string
}

type Event struct {
	Type     EventType
	Callback func(EventParams) error
}

func (f Feature) RunEvent(routePath string, params EventParams) error {
	if f.Event.Type == params.Type {
		return f.Event.Callback(params)
	}
	return errors.New("タイプが一致しません。")
}
