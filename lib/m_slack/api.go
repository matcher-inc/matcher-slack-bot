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

func PostResponse(params RequestParams, blocks Blocks) (err error) {
	options := blocks.toMsgOption(params)
	_, _, _, err = client.SendMessage("", slack.MsgOptionReplaceOriginal(params.responseURL), options)
	return
}

func DeleteOriginal(params RequestParams) (err error) {
	_, _, _, err = client.SendMessage("", slack.MsgOptionDeleteOriginal(params.responseURL))
	return
}

// モーダルなどを開く
// TriggerIDはとりあえず、後でparamsに含める
func OpenView(params RequestParams, modal Modal, TriggerID string) (err error) {
	_, err = client.OpenView(TriggerID, modal.ToViewRequest(params))
	return
}
