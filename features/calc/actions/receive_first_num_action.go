package actions

import (
	"fmt"
	mSlack "go-bot-test/lib/m_slack"
)

func receiveFirstNumCallback(params mSlack.RequestParams) error {
	if err := mSlack.DeleteOriginal(params); err != nil {
		return err
	}

	blocks := []mSlack.Block{
		mSlack.Text{
			Body: params.ActionParams.Value,
		},
		mSlack.Buttons{
			ActionPath: ReceiveOperatorAction.ActionPath,
			Options: []mSlack.ButtonsOption{
				{
					Label: "+",
					Value: fmt.Sprintf("%s +", params.ActionParams.Value),
				},
				{
					Label: "-",
					Value: fmt.Sprintf("%s -", params.ActionParams.Value),
				},
				{
					Label: "x",
					Value: fmt.Sprintf("%s x", params.ActionParams.Value),
				},
				{
					Label: "÷",
					Value: fmt.Sprintf("%s ÷", params.ActionParams.Value),
				},
			},
		},
	}

	if err := mSlack.PostPrivate(params, blocks); err != nil {
		return err
	}
	return nil
}
