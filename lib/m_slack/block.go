package mSlack

import (
	"github.com/slack-go/slack"
)

type Block interface {
	toBlock(RequestParams) slack.Block
}

type Blocks []Block
