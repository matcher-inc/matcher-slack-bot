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
	return slack.NewActionBlock(params.RequestKey+":"+s.ActionKey, selectMenu)
}

func (s Select) toBlockElement(params RequestParams) slack.BlockElement {
	return slack.NewOptionsSelectBlockElement(
		slack.OptTypeStatic,
		// NOTE: modalで使うとき、placeholderがあるとinvalid_arguments
		// slack.NewTextBlockObject(slack.PlainTextType, s.Placeholder, false, false)
		nil,
		params.RequestKey+":"+s.ActionKey,
		s.optionObjects(params)...,
	)
}
