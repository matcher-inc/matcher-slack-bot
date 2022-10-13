package mSlack

import (
	"go-bot-test/config/env"

	"github.com/slack-go/slack"
)

var (
	client = slack.New(env.SLACK_BOT_TOKEN)
)

func Post(params EventParams, blocks Blocks) (err error) {
	options := blocks.toMsgOption(params)

	_, _, err = client.PostMessage(params.ChannelID, options)
	return
}

func PostPrivate(params EventParams, blocks Blocks) (err error) {
	options := blocks.toMsgOption(params)

	_, err = client.PostEphemeral(params.ChannelID, params.UserID, options)
	return
}
