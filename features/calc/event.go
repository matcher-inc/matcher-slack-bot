package calc

import (
	"go-bot-test/features/calc/actions"
	"go-bot-test/lib/feature"
	mSlack "go-bot-test/lib/m_slack"
)

var event = feature.Event{
	Type:     mSlack.SlashEvent,
	Callback: eventCallback,
}

func eventCallback(params mSlack.RequestParams) error {
	blocks := []mSlack.Block{
		mSlack.Buttons{
			ActionPath: actions.ReceiveFirstNumAction.ActionPath,
			Options: []mSlack.ButtonsOption{
				{
					Label: "1",
					Value: "1",
				},
				{
					Label: "2",
					Value: "2",
				},
				{
					Label: "3",
					Value: "3",
				},
			},
		},
	}

	if err := mSlack.PostPrivate(params, blocks); err != nil {
		return err
	}
	return nil
}
