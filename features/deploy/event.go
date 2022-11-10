package deploy

import (
	"errors"
	"go-bot-test/features/deploy/actions"
	"go-bot-test/lib/feature"
	mSlack "go-bot-test/lib/m_slack"
	"log"
)

var event = feature.Event{
	Type:     mSlack.SlashEvent,
	Callback: eventCallback,
}

func eventCallback(params mSlack.RequestParams) error {
	blocks := []mSlack.Block{
		mSlack.Text{
			Body: "Please select *version*.",
		},
		mSlack.Select{
			ActionPath:  actions.SelectVersionAction.ActionPath,
			Placeholder: "Select version",
			Options: []mSlack.Option{
				{Label: "v1.0.0", Value: "v1.0.0", Description: "ドキドキの初バージョン"},
				{Label: "v1.1.0", Value: "v1.1.0", Description: "ワクワクのアップデート"},
				{Label: "v1.1.1", Value: "v1.1.1", Description: "気づくかな？マイナーアップデート"},
			},
		},
	}

	if err := mSlack.PostPrivate(params, blocks); err != nil {
		log.Println(err)
		return errors.New("エラー")
	}
	return nil
}
