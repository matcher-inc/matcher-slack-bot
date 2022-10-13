package mSlack

import (
	"github.com/slack-go/slack"
)

type RadioButtons struct {
	ActionKey string
	Options   []Option
}

func (r RadioButtons) optionObjects(params RequestParams) []*slack.OptionBlockObject {
	options := make([]*slack.OptionBlockObject, len(r.Options))
	for i, option := range r.Options {
		options[i] = option.toBlockObject(params)
	}
	return options
}

func (r RadioButtons) toBlock(params RequestParams) slack.Block {
	element := r.toBlockElement(params)
	// BlockIDを渡すから
	return slack.NewActionBlock(params.RequestKey+":"+r.ActionKey, element)
}

func (r RadioButtons) toBlockElement(params RequestParams) slack.BlockElement {
	return slack.NewRadioButtonsBlockElement(
		params.RequestKey+":"+r.ActionKey,
		r.optionObjects(params)...,
	)
}
