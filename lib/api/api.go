package api

import (
	"go-bot-test/config"

	"github.com/slack-go/slack"
)

var (
	Shared = slack.New(config.SLACK_BOT_TOKEN)
)
