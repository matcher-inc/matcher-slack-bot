package env

import "os"

var (
	SLACK_SIGNING_SECRET     = os.Getenv("SLACK_SIGNING_SECRET")
	SLACK_BOT_TOKEN          = os.Getenv("SLACK_BOT_TOKEN")
	SLACK_VERIFICATION_TOKEN = os.Getenv("SLACK_VERIFICATION_TOKEN")
)
