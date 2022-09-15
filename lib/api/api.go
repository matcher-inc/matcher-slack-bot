package api

import (
	"encoding/json"
	"errors"
	"go-bot-test/config/env"
	"net/http"

	"github.com/slack-go/slack"
)

var (
	Shared = slack.New(env.SLACK_BOT_TOKEN)
)

// 下記コードの動かし方がわからないため直接レスポンスに書き込む（invalid_trigger_id と怒られる）
// api.Shared.PushView(payload.TriggerID, *modal)
func PushModalView(modal *slack.ModalViewRequest, w http.ResponseWriter) error {
	resAction := slack.NewPushViewSubmissionResponse(modal)
	rBytes, err := json.Marshal(resAction)
	if err != nil {
		return errors.New("エラー")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rBytes)
	return nil
}
