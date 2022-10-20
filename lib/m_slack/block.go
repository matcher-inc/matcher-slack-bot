package mSlack

import (
	"github.com/slack-go/slack"
)

type Block interface {
	toOption(EventParams) slack.Block
}

type Blocks []Block

func (blocks Blocks) toMsgOption(params EventParams) slack.MsgOption {
	blockArr := make([]slack.Block, len(blocks))
	for i, b := range blocks {
		blockArr[i] = b.toOption(params)
	}
	return slack.MsgOptionBlocks(blockArr...)
}
