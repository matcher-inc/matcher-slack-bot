package web

import (
	"go-bot-test/config/routes"
	mSlack "go-bot-test/lib/m_slack"
	"net/http"
)

func handleEvent(w http.ResponseWriter, r *http.Request) {
	params, requestBody, eventType, err := mSlack.ParseEvent(r)
	if err != nil {
		raiseError(w, err)
		return
	}

	if eventType == mSlack.URLVerification {
		mSlack.VerificateUrl(w, requestBody)
		return
	}

	route, err := routes.GetRoute(params.FeaturePath)
	if err != nil {
		raiseError(w, err)
		return
	}

	route.Feature.RunEvent(*params, eventType)
}
