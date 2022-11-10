package actions

import (
	"encoding/json"
	"errors"
	mSlack "go-bot-test/lib/m_slack"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

var (
	reqButtonPushedAction          = "buttonPushedAction"
	reqOrderModalSubmission        = "orderModalSubmission"
	reqConfirmationModalSubmission = "confirmationModalSubmission"
	reqUnknown                     = "unknown"
)

func showDialogCallback(routePath string, payload slack.InteractionCallback, w http.ResponseWriter) error {
	modal := createOrderModalBySDK()

	modal.CallbackID = routePath + ":" + ReceiveFormAction.ActionPath
	modal.ExternalID = payload.User.ID + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	params := privateMeta{
		ChannelID: payload.Channel.ID,
	}
	bytes, err := json.Marshal(params)
	if err != nil {
		return errors.New("エラー")
	}
	modal.PrivateMetadata = string(bytes)

	if err := mSlack.OpenView(mSlack.RequestParams{}, *modal, payload.TriggerID); err != nil {
		log.Println(err)
		return errors.New("エラー")
	}
	return nil
}

func createOrderModalBySDK() *mSlack.Modal {
	shopTextSection := mSlack.Text{
		Body: ":hamburger: *Hey! Thank you for choosing us! We'll promise you to be full.*",
		Type: mSlack.MarkdownType,
	}
	dividerBlock := mSlack.Divider{}

	menuInput := mSlack.Input{
		BlockID: "block_id_menu",
		Label:   "Which one you want to have?",
		Element: mSlack.RadioButtons{
			ActionKey: "action_id_menu",
			Options: []mSlack.Option{
				{Label: "hamburger", Value: "hamburger"},
				{Label: "cheese_burger", Value: "cheese_burger"},
				{Label: "blt_burger", Value: "blt_burger"},
				{Label: "big_burger", Value: "big_burger"},
				{Label: "king_burger", Value: "king_burger"},
			},
		},
	}

	steakInput := mSlack.Input{
		BlockID: "action_id_steak",
		Label:   "How do you like your steak?",
		Element: mSlack.Select{
			ActionKey: "action_id_steak", // 不要？,
			Options: []mSlack.Option{
				{Label: "well done", Value: "well done"},
				{Label: "medium", Value: "medium"},
				{Label: "rare", Value: "rare"},
				{Label: "blue", Value: "blue"},
			},
		},
	}

	noteInput := mSlack.Input{
		BlockID: "block_id_note",
		Label:   "Anything else you want to tell us?",
		Element: mSlack.PlainText{
			ActionKey: "block_id_note",
			Multiline: true,
		},
	}

	modal := mSlack.Modal{
		Title:        "Hungryman Hamburgers",
		CloseButton:  "キャンセル",
		SubmitButton: "送信",
		Blocks: []mSlack.Block{
			shopTextSection,
			dividerBlock,
			menuInput,
			steakInput,
			noteInput,
		},
	}

	return &modal
}
