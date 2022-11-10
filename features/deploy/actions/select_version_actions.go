package actions

import (
	"fmt"
	mSlack "go-bot-test/lib/m_slack"
)

func selectVersionActionCallback(params mSlack.RequestParams) error {
	blocks := []mSlack.Block{
		mSlack.Text{Body: fmt.Sprintf("Could I deploy `%s`?", params.ActionParams.Value)},
		mSlack.Buttons{
			ActionPath: ConfirmDeploymentAction.ActionPath,
			Options: []mSlack.ButtonsOption{
				{
					Label: "Do it",
					Value: params.ActionParams.Value,
					Type:  mSlack.ButtonTypePrimary,
				},
				{
					Label: "Stop",
					Value: "deny",
					Type:  mSlack.ButtonTypeDanger,
				},
			},
		},
	}
	if err := mSlack.PostPrivate(params, blocks); err != nil {
		return err
	}
	return nil
}
