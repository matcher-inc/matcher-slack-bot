package mSlack

import (
	"github.com/slack-go/slack"
)

type Divider struct {
}

func (t Divider) toOption(_ EventParams) slack.Block {
	return slack.NewDividerBlock()
}
