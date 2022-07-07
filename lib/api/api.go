package api

import (
	"go-bot-test/config/env"

	"github.com/slack-go/slack"
)

var (
	Shared = slack.New(env.SLACK_BOT_TOKEN)
)
