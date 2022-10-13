package mSlack

import (
	"github.com/slack-go/slack"
)

type Divider struct {
}

func (t Divider) toBlock(_ RequestParams) slack.Block {
	return slack.NewDividerBlock()
}
