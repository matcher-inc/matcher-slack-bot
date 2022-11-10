package actions

import (
	"encoding/json"
	"errors"
	"go-bot-test/lib/api"
	mSlack "go-bot-test/lib/m_slack"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

func receiveFormCallback(params mSlack.RequestParams) error {
	// Get the selected information.
	// - radio button
	menu := params.ActionParams.Values["block_id_menu"]["action_id_menu"].SelectedOption.Value

	// - static_select
	steak := params.ActionParams.Values["block_id_steak"]["action_id_steak"].SelectedOption.Value

	// - text
	note := params.ActionParams.Values["block_id_note"]["action_id_note"].Value

	// Create a confirmation modal.
	// - apperance
	modal := createConfirmationModalBySDK(menu, steak, note)

	// - metadata : CallbackID
	modal.CallbackID = params.FeaturePath + ":" + ConfirmAction.ActionPath

	// - metadata : ExternalID
	modal.ExternalID = params.UserID + strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	// - metadata : PrivateMeta
	//   - Get private metadata of a message
	var pMeta privateMeta
	if err := json.Unmarshal([]byte(params.PrivateMetadata), &pMeta); err != nil {
		// return events.APIGatewayProxyResponse{StatusCode: 200}, fmt.Errorf("failed to unmarshal private metadata: %w", err)
		return errors.New("エラー")
	}

	//   - Create new private metadata
	privateParams := privateMeta{
		ChannelID: pMeta.ChannelID,
		order: order{
			Menu:   menu,
			Steak:  steak,
			Note:   note,
			Amount: "700",
		},
	}

	pBytes, err := json.Marshal(privateParams)
	if err != nil {
		return errors.New("エラー")
	}

	modal.PrivateMetadata = string(pBytes)

	// api.PushModalView(modal, w)
	// return nil
	// 上記もしくは下記

	// payload.View.ExternalID, payload.View.ID のどちらかだけを渡す
	// 両方渡すとargumentserror
	if _, err := api.Shared.UpdateView(*modal, params.ExternalID, "", ""); err != nil {
		return errors.New("エラー")
	}
	return nil
}

func createConfirmationModalBySDK(menu, steak, note string) *slack.ModalViewRequest {

	// Create a modal.
	// - Text section
	titleText := slack.NewTextBlockObject("mrkdwn", ":wave: *🐷確認するよ confirmation*", false, false)
	titleTextSection := slack.NewSectionBlock(titleText, nil, nil)

	// Divider
	dividerBlock := slack.NewDividerBlock()

	// - Text section
	sMenuText := slack.NewTextBlockObject("mrkdwn", "*Menu :hamburger:*\n ... "+menu, false, false)
	sMenuTextSection := slack.NewSectionBlock(sMenuText, nil, nil)

	// - Text section
	sSteakText := slack.NewTextBlockObject("mrkdwn", "*How do you like your steak?*\n"+steak, false, false)
	sSteakTextSection := slack.NewSectionBlock(sSteakText, nil, nil)

	// - Text section
	sNoteText := slack.NewTextBlockObject("mrkdwn", "*Anything else you want to tell us?*\n"+note, false, false)
	sNoteTextSection := slack.NewSectionBlock(sNoteText, nil, nil)

	// - Text section
	amountText := slack.NewTextBlockObject("mrkdwn", "*Amount :moneybag:*\n$ 700", false, false)
	amountTextSection := slack.NewSectionBlock(amountText, nil, nil)

	// - Input with plain_text_input
	chipText := slack.NewTextBlockObject("plain_text", "Chip ($)", false, false)
	chipInputElement := slack.NewPlainTextInputBlockElement(nil, "action_id_chip")
	chipInput := slack.NewInputBlock("block_id_chip", chipText, chipInputElement)
	chipHintText := slack.NewTextBlockObject("plain_text", "Thank you for your kindness!", false, false)
	chipInput.Hint = chipHintText
	chipInput.Optional = true

	// Blocks
	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			titleTextSection,
			dividerBlock,
			sMenuTextSection,
			sSteakTextSection,
			sNoteTextSection,
			dividerBlock,
			amountTextSection,
			chipInput,
		},
	}

	// ModalView
	modal := slack.ModalViewRequest{
		Type:   slack.ViewType("modal"),
		Title:  slack.NewTextBlockObject("plain_text", "Hungryman Hamburgers", false, false),
		Close:  slack.NewTextBlockObject("plain_text", "Cancel", false, false),
		Submit: slack.NewTextBlockObject("plain_text", "Order!", false, false),
		Blocks: blocks,
	}

	return &modal
}

type privateMeta struct {
	ChannelID string `json:"channel_id"`
	order
}

type order struct {
	Menu   string `json:"order_menu"`
	Steak  string `json:"order_steak"`
	Note   string `json:"order_note"`
	Amount string `json:"order_amount"`
}
