package mSlack

import (
	"github.com/slack-go/slack"
)

type InputElement interface {
	toBlockElement(EventParams) slack.BlockElement
}

type Input struct {
	BlockID string
	Label   string
	Element InputElement
}

func (i Input) toOption(params EventParams) slack.Block {
	label := slack.NewTextBlockObject(slack.MarkdownType, i.Label, false, false)
	return slack.NewInputBlock(i.BlockID, label, i.Element.toBlockElement(params))
}
