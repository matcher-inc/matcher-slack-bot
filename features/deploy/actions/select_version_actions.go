package actions

import (
	"errors"
	"fmt"
	"go-bot-test/app/constants"
	"go-bot-test/lib/api"
	"log"

	"github.com/slack-go/slack"
)

func SelectVersionActionCallback(payload slack.InteractionCallback) error {
	action := payload.ActionCallback.BlockActions[0]
	version := action.SelectedOption.Value

	text := slack.NewTextBlockObject(slack.MarkdownType,
		fmt.Sprintf("Could I deploy `%s`?", version), false, false)
	textSection := slack.NewSectionBlock(text, nil, nil)

	confirmButtonText := slack.NewTextBlockObject(slack.PlainTextType, "Do it", false, false)
	confirmButton := slack.NewButtonBlockElement("", version, confirmButtonText)
	confirmButton.WithStyle(slack.StylePrimary)

	denyButtonText := slack.NewTextBlockObject(slack.PlainTextType, "Stop", false, false)
	denyButton := slack.NewButtonBlockElement("", "deny", denyButtonText)
	denyButton.WithStyle(slack.StyleDanger)

	actionBlock := slack.NewActionBlock(constants.ConfirmDeploymentAction, confirmButton, denyButton)

	fallbackText := slack.MsgOptionText("This client is not supported.", false)
	blocks := slack.MsgOptionBlocks(textSection, actionBlock)

	replaceOriginal := slack.MsgOptionReplaceOriginal(payload.ResponseURL)
	if _, _, _, err := api.Shared.SendMessage("", replaceOriginal, fallbackText, blocks); err != nil {
		log.Println(err)
		return errors.New("エラー")
	}
	return errors.New("エラー")
}
