package actions

import (
	"fmt"
	mSlack "go-bot-test/lib/m_slack"
	"strings"
)

func confirmDeploymentActionCallback(params mSlack.RequestParams) error {
	if err := mSlack.DeleteOriginal(params); err != nil {
		return err
	}
	if strings.HasPrefix(params.ActionParams.Value, "v") {
		endMsgBlocks := []mSlack.Block{
			mSlack.Text{Body: fmt.Sprintf("`%s` deployment completed!", params.ActionParams.Value)},
		}
		if err := mSlack.PostPrivate(params, endMsgBlocks); err != nil {
			return err
		}
	}
	return nil
}
