package mSlack

import (
	"github.com/slack-go/slack"
)

type Block interface {
	toBlock(EventParams) slack.Block
}

type Blocks []Block
