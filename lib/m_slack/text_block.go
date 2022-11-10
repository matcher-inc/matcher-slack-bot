package mSlack

import (
	"github.com/slack-go/slack"
)

const (
	MarkdownType    = "mrkdwn"
	PlainTextType   = "plain_text"
	motConfirmation = "confirm"
	motOption       = "option"
	motOptionGroup  = "option_group"
)

type Text struct {
	Body string
	Type string
}

func (t Text) toBlock(params RequestParams) slack.Block {
	return slack.NewSectionBlock(t.toBlockObject(params), nil, nil)
}

func (t Text) toBlockObject(para RequestParams) *slack.TextBlockObject {
	// NOTE: modalで使うとき、slack.MarkdownTypeだとinvalid_arguments
	// return slack.NewTextBlockObject(slack.MarkdownType, t.Body, false, false)
	return slack.NewTextBlockObject(t.textType(), t.Body, false, false)
}

func (t Text) textType() string {
	if len(t.Type) > 0 {
		return t.Type
	}
	return PlainTextType
}
