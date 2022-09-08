package mSlack

import (
	"github.com/slack-go/slack"
)

type Text struct {
	Body string
}

func (t Text) toOption(_ EventParams) slack.Block {
	textObject := slack.NewTextBlockObject(slack.MarkdownType, t.Body, false, false)
	return slack.NewSectionBlock(textObject, nil, nil)
}
