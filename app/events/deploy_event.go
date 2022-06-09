package events

import (
	"go-bot-test/app/constants"
	"go-bot-test/lib/api"
	"go-bot-test/lib/listner"
	"log"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func DeployEvent(event slackevents.AppMentionEvent) (error listner.EventError) {
	text := slack.NewTextBlockObject(slack.MarkdownType, "Please select *version*.", false, false)
	textSection := slack.NewSectionBlock(text, nil, nil)

	versions := []string{"v1.0.0", "v1.1.0", "v1.1.1"}
	options := make([]*slack.OptionBlockObject, 0, len(versions))
	for _, v := range versions {
		optionText := slack.NewTextBlockObject(slack.PlainTextType, v, false, false)
		options = append(options, slack.NewOptionBlockObject(v, optionText, optionText))
	}

	placeholder := slack.NewTextBlockObject(slack.PlainTextType, "Select version", false, false)
	selectMenu := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, placeholder, "", options...)

	actionBlock := slack.NewActionBlock(constants.SelectVersionAction, selectMenu)

	fallbackText := slack.MsgOptionText("This client is not supported.", false)
	blocks := slack.MsgOptionBlocks(textSection, actionBlock)

	if _, err := api.Shared.PostEphemeral(event.Channel, event.User, fallbackText, blocks); err != nil {
		log.Println(err)
		return listner.EventStandardError
	}
	return
}
