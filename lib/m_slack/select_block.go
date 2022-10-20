package mSlack

import (
	"github.com/slack-go/slack"
)

type Select struct {
	ActionKey   string
	Placeholder string
	Options     []Option
}

func (s Select) optionObjects(params RequestParams) []*slack.OptionBlockObject {
	options := make([]*slack.OptionBlockObject, len(s.Options))
	for i, option := range s.Options {
		options[i] = option.toBlockObject(params)
	}
	return options
}

func (s Select) toBlock(params RequestParams) slack.Block {
	selectMenu := s.toBlockElement(params)
	return slack.NewActionBlock(params.FeaturePath+":"+s.ActionKey, selectMenu)
}

func (s Select) toBlockElement(params RequestParams) slack.BlockElement {
	placeholder := slack.NewTextBlockObject(slack.PlainTextType, s.Placeholder, false, false)
	return slack.NewOptionsSelectBlockElement(
		slack.OptTypeStatic,
		placeholder,
		params.FeaturePath+":"+s.ActionKey,
		s.optionObjects(params)...,
	)
}
