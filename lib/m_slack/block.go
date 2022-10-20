package mSlack

import (
	"github.com/slack-go/slack"
)

type Block interface {
	toBlock(RequestParams) slack.Block
}

type Blocks []Block

func (blocks Blocks) toMsgOption(params RequestParams) slack.MsgOption {
	blockArr := make([]slack.Block, len(blocks))
	for i, b := range blocks {
		blockArr[i] = b.toBlock(params)
	}
	return slack.MsgOptionBlocks(blockArr...)
}
