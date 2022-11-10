package feature

import (
	"errors"
	mSlack "go-bot-test/lib/m_slack"
)

type Action struct {
	ActionPath string
	Callback   func(mSlack.RequestParams) error
}

func (f Feature) RunAction(params mSlack.RequestParams) error {
	for _, action := range f.Actions {
		if action.ActionPath == params.ActionPath {
			return action.Callback(params)
		}
	}
	return errors.New("タイプが一致しません。")
}
