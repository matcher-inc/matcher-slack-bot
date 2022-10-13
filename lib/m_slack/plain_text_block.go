package mSlack

import (
	"github.com/slack-go/slack"
)

type PlainText struct {
	Multiline bool
	ActionKey string
}

func (p PlainText) toBlock(params EventParams) slack.Block {
	element := p.toBlockElement(params)
	return slack.NewActionBlock(params.RequestKey+":"+p.ActionKey, element)
}

func (p PlainText) toBlockElement(params EventParams) slack.BlockElement {
	el := slack.NewPlainTextInputBlockElement(nil, params.RequestKey+":"+p.ActionKey)
	el.Multiline = p.Multiline
	return el
}
