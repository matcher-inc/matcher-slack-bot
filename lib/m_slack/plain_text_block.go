package mSlack

import (
	"github.com/slack-go/slack"
)

type PlainText struct {
	Multiline bool
	ActionKey string
}

func (p PlainText) toBlock(params RequestParams) slack.Block {
	element := p.toBlockElement(params)
	return slack.NewActionBlock(params.FeaturePath+":"+p.ActionKey, element)
}

func (p PlainText) toBlockElement(params RequestParams) slack.BlockElement {
	el := slack.NewPlainTextInputBlockElement(nil, params.FeaturePath+":"+p.ActionKey)
	el.Multiline = p.Multiline
	return el
}
