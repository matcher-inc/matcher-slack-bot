package mSlack

import (
	"github.com/slack-go/slack"
)

type RadioButtons struct {
	ActionKey string
	Options   []Option
}

func (r RadioButtons) optionObjects(params EventParams) []*slack.OptionBlockObject {
	options := make([]*slack.OptionBlockObject, len(r.Options))
	for i, option := range r.Options {
		options[i] = option.toBlockObject(params)
	}
	return options
}

func (r RadioButtons) toBlock(params EventParams) slack.Block {
	element := r.toBlockElement(params)
	// BlockIDを渡すから
	return slack.NewActionBlock(params.RequestKey+":"+r.ActionKey, element)
}

func (r RadioButtons) toBlockElement(params EventParams) slack.BlockElement {
	return slack.NewRadioButtonsBlockElement(
		params.RequestKey+":"+r.ActionKey,
		r.optionObjects(params)...,
	)
}
