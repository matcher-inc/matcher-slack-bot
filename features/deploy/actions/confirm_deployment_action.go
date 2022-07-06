package actions

import (
	"errors"
	"fmt"
	"go-bot-test/lib/api"
	"log"
	"strings"
	"time"

	"github.com/slack-go/slack"
)

func ConfirmDeploymentAction(payload slack.InteractionCallback) error {
	action := payload.ActionCallback.BlockActions[0]
	if strings.HasPrefix(action.Value, "v") {
		version := action.Value
		go func() {
			startMsg := slack.MsgOptionText(
				fmt.Sprintf("<@%s> OK, I will deploy `%s`.", payload.User.ID, version), false)
			if _, _, err := api.Shared.PostMessage(payload.Channel.ID, startMsg); err != nil {
				log.Println(err)
			}

			deploy(version)

			endMsg := slack.MsgOptionText(
				fmt.Sprintf("`%s` deployment completed!", version), false)
			if _, _, err := api.Shared.PostMessage(payload.Channel.ID, endMsg); err != nil {
				log.Println(err)
			}
		}()
	}

	deleteOriginal := slack.MsgOptionDeleteOriginal(payload.ResponseURL)
	if _, _, _, err := api.Shared.SendMessage("", deleteOriginal); err != nil {
		log.Println(err)
		return errors.New("エラー")
	}
	return errors.New("エラー")
}

// 追加
func deploy(version string) {
	log.Printf("deploy %s", version)
	time.Sleep(10 * time.Second)
}
