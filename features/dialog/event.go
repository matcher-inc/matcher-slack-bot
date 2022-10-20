package dialog

import (
	"errors"
	"go-bot-test/features/dialog/actions"
	"go-bot-test/lib/feature"
	mSlack "go-bot-test/lib/m_slack"
	"log"
)

var event = feature.Event{
	Type:     mSlack.AppMentionEvent,
	Callback: eventCallback,
}

func eventCallback(params mSlack.RequestParams) error {
	blocks := []mSlack.Block{
		mSlack.Text{
			Body: "Dialog Event Text Body",
		},
		mSlack.Button{
			ActionKey: actions.ShowDialogAction.ActionPath,
			Text:      "Dialog Event Button Text",
			Type:      mSlack.ButtonTypePrimary,
		},
	}
	if err := mSlack.Post(params, blocks); err != nil {
		log.Printf("[ERROR] Failed to send a message to Slack: %v", err)
		return errors.New("エラー")
	}
	return nil
}
