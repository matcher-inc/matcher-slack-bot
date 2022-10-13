package feature

import (
	"errors"
	mSlack "go-bot-test/lib/m_slack"
)

type Event struct {
	Type     mSlack.EventType
	Callback func(mSlack.RequestParams) error
}

func (f Feature) RunEvent(params mSlack.RequestParams) error {
	if f.Event.Type == params.Type {
		return f.Event.Callback(params)
	}
	return errors.New("タイプが一致しません。")
}
