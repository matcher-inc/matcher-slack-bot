package mSlack

import (
	"github.com/slack-go/slack"
)

type Block interface {
	toBlock(RequestParams) slack.Block
}

type Blocks []Block

func (blocks Blocks) toBlockArray(params RequestParams) []slack.Block {
	blockArr := make([]slack.Block, len(blocks))
	for i, b := range blocks {
		blockArr[i] = b.toBlock(params)
	}
	return blockArr
}

func (blocks Blocks) toMsgOption(params RequestParams) slack.MsgOption {
	return slack.MsgOptionBlocks(blocks.toBlockArray(params)...)
}
