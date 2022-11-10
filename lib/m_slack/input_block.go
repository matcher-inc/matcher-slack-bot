package mSlack

import (
	"github.com/slack-go/slack"
)

type InputElement interface {
	toBlockElement(RequestParams) slack.BlockElement
}

type Input struct {
	BlockID string
	Label   string
	Element InputElement
}

func (i Input) toBlock(params RequestParams) slack.Block {
	// NOTE: modalで使うとき、slack.MarkdownTypeだとinvalid_arguments
	// label := slack.NewTextBlockObject(slack.MarkdownType, i.Label, false, false)
	label := slack.NewTextBlockObject(slack.PlainTextType, i.Label, false, false)
	return slack.NewInputBlock(i.BlockID, label, i.Element.toBlockElement(params))
}