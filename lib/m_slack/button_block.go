package mSlack

import (
	"github.com/slack-go/slack"
)

type buttonType struct {
	value slack.Style
}

type Button struct {
	ActionKey string
	Text      string
	Type      buttonType
}

var (
	ButtonTypePrimary = buttonType{value: slack.StylePrimary}
	ButtonTypeDanger  = buttonType{value: slack.StyleDanger}
)

func (b Button) toBlock(params RequestParams) slack.Block {
	confirmButtonText := slack.NewTextBlockObject(slack.PlainTextType, b.Text, false, false)
	confirmButton := slack.NewButtonBlockElement("", "", confirmButtonText)
	confirmButton.WithStyle(b.Type.value)
	return slack.NewActionBlock(params.FeaturePath+":"+b.ActionKey, confirmButton)
}
