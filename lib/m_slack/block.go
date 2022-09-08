package mSlack

import (
	"github.com/slack-go/slack"
)

type Block interface {
	toOption(EventParams) slack.Block
}

type Blocks []Block
