package mSlack

import (
	"go-bot-test/config/env"

	"github.com/slack-go/slack"
)

var (
	client = slack.New(env.SLACK_BOT_TOKEN)
)

// リクエストをしたチャンネルに送信
func Post(params RequestParams, blocks Blocks) (err error) {
	options := blocks.toMsgOption(params)
	_, _, err = client.PostMessage(params.ChannelID, options)
	return
}

// リクエストをしたチャンネルに、リクエストをしたユーザーにだけ見えるメッセージを送信
func PostPrivate(params RequestParams, blocks Blocks) (err error) {
	options := blocks.toMsgOption(params)
	_, err = client.PostEphemeral(params.ChannelID, params.UserID, options)
	return
}

// リクエストをしたユーザーのダイレクトメッセージに送信
func PostDirect(params RequestParams, blocks Blocks) (err error) {
	options := blocks.toMsgOption(params)
	_, _, err = client.PostMessage(params.UserID, options)
	return
}
