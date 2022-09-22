package mSlack

import (
	"go-bot-test/config/env"

	"github.com/slack-go/slack"
)

var (
	client = slack.New(env.SLACK_BOT_TOKEN)
)

func PostPrivate(params EventParams, blocks Blocks) (err error) {
	blockArr := make([]slack.Block, len(blocks))
	for i, b := range blocks {
		blockArr[i] = b.toOption(params)
	}
	options := slack.MsgOptionBlocks(blockArr...)
	_, err = client.PostEphemeral(params.ChannelID, params.UserID, options)
	return
}
