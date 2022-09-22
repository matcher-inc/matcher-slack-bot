package mSlack

import (
	"github.com/slack-go/slack"
)

type Option struct {
	Label       string
	Value       string
	Description string
}

type Select struct {
	ActionKey   string
	Placeholder string
	Options     []Option
}

func (s Select) toOption(params EventParams) slack.Block {
	options := make([]*slack.OptionBlockObject, len(s.Options))
	for i, option := range s.Options {
		label := slack.NewTextBlockObject(slack.PlainTextType, option.Label, false, false)
		descriptionText := option.Description
		if descriptionText == "" {
			descriptionText = option.Label
		}
		description := slack.NewTextBlockObject(slack.PlainTextType, descriptionText, false, false)
		options[i] = slack.NewOptionBlockObject(option.Value, label, description)
	}
	placeholder := slack.NewTextBlockObject(slack.PlainTextType, s.Placeholder, false, false)
	selectMenu := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, placeholder, "", options...)

	return slack.NewActionBlock(params.RequestKey+":"+s.ActionKey, selectMenu)
}
