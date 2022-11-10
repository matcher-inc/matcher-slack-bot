package mSlack

import (
	"github.com/slack-go/slack"
)

type buttonType struct {
	value slack.Style
}

type Button struct {
	ActionPath string
	Value      string
	Text       string
	Type       buttonType
}

type Buttons struct {
	ActionPath string
	Options    []ButtonsOption
}

type ButtonsOption struct {
	Label string
	Value string
	Type  buttonType
}

var (
	ButtonTypePrimary = buttonType{value: slack.StylePrimary}
	ButtonTypeDanger  = buttonType{value: slack.StyleDanger}
)

func (b Button) toBlock(params RequestParams) slack.Block {
	buttonText := slack.NewTextBlockObject(slack.PlainTextType, b.Text, false, false)
	button := slack.NewButtonBlockElement("", b.Value, buttonText)
	button.WithStyle(b.Type.value)
	return slack.NewActionBlock(params.FeaturePath+":"+b.ActionPath, button)
}

func (o ButtonsOption) toOptionBlock() slack.BlockElement {
	buttonText := slack.NewTextBlockObject(slack.PlainTextType, o.Label, false, false)
	button := slack.NewButtonBlockElement("", o.Value, buttonText)
	button.WithStyle(o.Type.value)
	return *button
}

func (bs Buttons) toBlock(params RequestParams) slack.Block {
	options := make([]slack.BlockElement, len(bs.Options))
	for i, option := range bs.Options {
		options[i] = option.toOptionBlock()
	}
	return slack.NewActionBlock(params.FeaturePath+":"+bs.ActionPath, options...)
}
