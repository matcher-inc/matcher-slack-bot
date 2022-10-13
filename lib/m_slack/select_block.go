package mSlack

import (
	"github.com/slack-go/slack"
)

type Select struct {
	ActionKey   string
	Placeholder string
	Options     []Option
}

func (s Select) optionObjects(params EventParams) []*slack.OptionBlockObject {
	options := make([]*slack.OptionBlockObject, len(s.Options))
	for i, option := range s.Options {
		options[i] = option.toBlockObject(params)
	}
	return options
}

func (s Select) toOption(params EventParams) slack.Block {
	selectMenu := s.toBlockElement(params)
	return slack.NewActionBlock(params.RequestKey+":"+s.ActionKey, selectMenu)
}

func (s Select) toBlockElement(params EventParams) slack.BlockElement {
	placeholder := slack.NewTextBlockObject(slack.PlainTextType, s.Placeholder, false, false)
	return slack.NewOptionsSelectBlockElement(
		slack.OptTypeStatic,
		placeholder,
		params.RequestKey+":"+s.ActionKey,
		s.optionObjects(params)...,
	)
}
