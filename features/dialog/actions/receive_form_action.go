package actions

import (
	"encoding/json"
	"errors"
	mSlack "go-bot-test/lib/m_slack"
	"strconv"
	"time"
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
	if err := mSlack.UpdateView(params, *modal); err != nil {
		return errors.New("エラー")
	}
	return nil
}

func createConfirmationModalBySDK(menu, steak, note string) *mSlack.Modal {

	dividerBlock := mSlack.Divider{}
	sMenuTextSection := mSlack.Text{
		Body: "*Menu :hamburger:*\n ... " + menu,
	}

	sSteakTextSection := mSlack.Text{
		Body: "*How do you like your steak?*\n" + steak,
	}
	sNoteTextSection := mSlack.Text{
		Body: "*Anything else you want to tell us?*\n" + note,
	}
	amountTextSection := mSlack.Text{
		Body: "*Amount :moneybag:*\n$ 700",
	}

	chipInput := mSlack.Input{
		BlockID: "action_id_chip",
		Label:   "Which one you want to have?",
		Element: mSlack.PlainText{
			ActionKey: "action_id_chip",
			Multiline: false,
		},
		Hint:     "Thank you for your kindness!",
		Optional: true,
	}

	modal := mSlack.Modal{
		Title:        "🐷確認するよ confirmation*",
		CloseButton:  "キャンセル",
		SubmitButton: "送信",
		Blocks: []mSlack.Block{
			dividerBlock,
			sMenuTextSection,
			sSteakTextSection,
			sNoteTextSection,
			dividerBlock,
			amountTextSection,
			chipInput,
		},
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
