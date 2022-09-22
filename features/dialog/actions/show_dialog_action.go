package actions

import (
	"encoding/json"
	"errors"
	"go-bot-test/lib/api"
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

	// - metadata : CallbackID
	modal.CallbackID = routePath + ":" + ReceiveFormAction.Key

	// - metadata : ExternalID
	modal.ExternalID = payload.User.ID + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	params := privateMeta{
		ChannelID: payload.Channel.ID,
	}
	bytes, err := json.Marshal(params)
	if err != nil {
		return errors.New("エラー")
	}
	modal.PrivateMetadata = string(bytes)

	if _, err := api.Shared.OpenView(payload.TriggerID, *modal); err != nil {
		log.Println(err)
		return errors.New("エラー")
	}
	return nil
}

func createOrderModalBySDK() *slack.ModalViewRequest {
	// Text section
	shopText := slack.NewTextBlockObject("mrkdwn", ":hamburger: *Hey! Thank you for choosing us! We'll promise you to be full.*", false, false)
	shopTextSection := slack.NewSectionBlock(shopText, nil, nil)

	// Divider
	dividerBlock := slack.NewDividerBlock()

	// Input with radio buttons
	optHamburgerText := slack.NewTextBlockObject("plain_text", "hamburger" /*"Hamburger"*/, false, false)
	optHamburgerObj := slack.NewOptionBlockObject("hamburger", optHamburgerText, optHamburgerText)

	optCheeseText := slack.NewTextBlockObject("plain_text", "cheese_burger" /*"Cheese burger"*/, false, false)
	optCheeseObj := slack.NewOptionBlockObject("cheese_burger", optCheeseText, optCheeseText)

	optBLTText := slack.NewTextBlockObject("plain_text", "blt_burger" /*"BLT burger"*/, false, false)
	optBLTObj := slack.NewOptionBlockObject("blt_burger", optBLTText, optBLTText)

	optBigText := slack.NewTextBlockObject("plain_text", "big_burger" /*"Big burger"*/, false, false)
	optBigObj := slack.NewOptionBlockObject("big_burger", optBigText, optBigText)

	optKingText := slack.NewTextBlockObject("plain_text", "king_burger" /*"King burger"*/, false, false)
	optKingObj := slack.NewOptionBlockObject("king_burger", optKingText, optKingText)

	menuElement := slack.NewRadioButtonsBlockElement("action_id_menu", optHamburgerObj, optCheeseObj, optBLTObj, optBigObj, optKingObj)

	menuLabel := slack.NewTextBlockObject("plain_text", "Which one you want to have?", false, false)
	menuInput := slack.NewInputBlock("block_id_menu", menuLabel, menuElement)

	// Input with static_select
	optWellDoneText := slack.NewTextBlockObject("plain_text", "well done", false, false)
	optWellDoneObj := slack.NewOptionBlockObject("well_done", optWellDoneText, optWellDoneText)

	optMediumText := slack.NewTextBlockObject("plain_text", "medium", false, false)
	optMediumObj := slack.NewOptionBlockObject("medium", optMediumText, optMediumText)

	optRareText := slack.NewTextBlockObject("plain_text", "rare", false, false)
	optRareObj := slack.NewOptionBlockObject("rare", optRareText, optRareText)

	optBlueText := slack.NewTextBlockObject("plain_text", "blue", false, false)
	optBlueObj := slack.NewOptionBlockObject("blue", optBlueText, optBlueText)

	steakInputElement := slack.NewOptionsSelectBlockElement("static_select", nil, "action_id_steak", optWellDoneObj, optMediumObj, optRareObj, optBlueObj)

	steakLabel := slack.NewTextBlockObject("plain_text", "How do you like your steak?", false, false)
	steakInput := slack.NewInputBlock("block_id_steak", steakLabel, steakInputElement)

	// Input with plain_text_input
	noteText := slack.NewTextBlockObject("plain_text", "Anything else you want to tell us?", false, false)
	noteInputElement := slack.NewPlainTextInputBlockElement(nil, "action_id_note")
	noteInputElement.Multiline = true
	noteInput := slack.NewInputBlock("block_id_note", noteText, noteInputElement)
	noteInput.Optional = true

	// Blocks
	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			shopTextSection,
			dividerBlock,
			menuInput,
			steakInput,
			noteInput,
		},
	}

	// ModalView
	modal := slack.ModalViewRequest{
		Type:   slack.ViewType("modal"),
		Title:  slack.NewTextBlockObject("plain_text", "Hungryman Hamburgers", false, false),
		Close:  slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Submit: slack.NewTextBlockObject("plain_text", "Submit", false, false),
		Blocks: blocks,
	}

	return &modal
}
