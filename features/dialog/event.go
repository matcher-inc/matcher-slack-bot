package dialog

import (
	"errors"
	"fmt"
	"go-bot-test/features/dialog/actions"
	"go-bot-test/lib/api"
	"go-bot-test/lib/feature"
	mSlack "go-bot-test/lib/m_slack"
	"log"

	"github.com/slack-go/slack"
)

var event = feature.Event{
	Type:     mSlack.SlashEvent,
	Callback: eventCallback,
}

func eventCallback(params mSlack.EventParams) error {
	// log.Printf("[ERROR] Failed to send a message to Slack: %v", err)
	list := createShopListBySDK(params.RequestKey)

	// Send a shop list to slack channel.
	if _, _, err := api.Shared.PostMessage(params.ChannelID, list); err != nil {
		log.Printf("[ERROR] Failed to send a message to Slack: %v", err)
		return errors.New("エラー")
	}
	return nil
}

func createShopListBySDK(requestKey string) slack.MsgOption {
	text := slack.NewTextBlockObject(slack.MarkdownType,
		fmt.Sprintf("Could I deploy `%s`?", ""), false, false)
	textSection := slack.NewSectionBlock(text, nil, nil)

	confirmButtonText := slack.NewTextBlockObject(slack.PlainTextType, "Do it", false, false)
	confirmButton := slack.NewButtonBlockElement("", "", confirmButtonText)
	confirmButton.WithStyle(slack.StylePrimary)

	actionBlock := slack.NewActionBlock(requestKey+":"+actions.ShowDialogAction.Key, confirmButton)

	blocks := slack.MsgOptionBlocks(textSection, actionBlock)
	return blocks
}
