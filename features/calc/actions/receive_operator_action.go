package actions

import (
	"fmt"
	mSlack "go-bot-test/lib/m_slack"
)

func receiveOperatorCallback(params mSlack.RequestParams) error {
	if err := mSlack.DeleteOriginal(params); err != nil {
		return err
	}

	blocks := []mSlack.Block{
		mSlack.Text{
			Body: params.ActionParams.Value,
		},
		mSlack.Buttons{
			ActionPath: CalcResultAction.ActionPath,
			Options: []mSlack.ButtonsOption{
				{
					Label: "1",
					Value: fmt.Sprintf("%s 1", params.ActionParams.Value),
				},
				{
					Label: "2",
					Value: fmt.Sprintf("%s 2", params.ActionParams.Value),
				},
				{
					Label: "3",
					Value: fmt.Sprintf("%s 3", params.ActionParams.Value),
				},
			},
		},
	}

	if err := mSlack.PostPrivate(params, blocks); err != nil {
		return err
	}
	return nil
}
