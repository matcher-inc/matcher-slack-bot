package mSlack

import (
	"github.com/slack-go/slack"
)

type Option struct {
	Label       string
	Value       string
	Description string
}

func (o Option) toBlockObject(params RequestParams) *slack.OptionBlockObject {
	label := slack.NewTextBlockObject(slack.PlainTextType, o.Label, false, false)
	descriptionText := o.Description
	if descriptionText == "" {
		descriptionText = o.Label
	}
	description := slack.NewTextBlockObject(slack.PlainTextType, descriptionText, false, false)
	return slack.NewOptionBlockObject(o.Value, label, description)
}
