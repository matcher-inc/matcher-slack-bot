package mSlack

import (
	"encoding/json"
	"errors"
	"go-bot-test/config/env"
	"go-bot-test/lib/feature"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func VerificateSigningSecret(r *http.Request) (slack.SecretsVerifier, error) {
	return slack.NewSecretsVerifier(r.Header, env.SLACK_SIGNING_SECRET)
}

func VerificateUrl(w http.ResponseWriter, params feature.EventParams) (err error) {
	var res *slackevents.ChallengeResponse

	err = json.Unmarshal(params.RequestBody, &res)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte(res.Challenge))
	if err != nil {
		return
	}
	return
}

func VerificateSlashToken(token string) error {
	if token != env.SLACK_VERIFICATION_TOKEN {
		return errors.New("Failed Verification Slash Token")
	}
	return nil
}