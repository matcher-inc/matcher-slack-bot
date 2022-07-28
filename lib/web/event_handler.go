package web

import (
	"go-bot-test/config/routes"
	"go-bot-test/lib/feature"
	mSlack "go-bot-test/lib/m_slack"
	"net/http"
)

func handleEvent(w http.ResponseWriter, r *http.Request) {
	params, err := mSlack.ParseEvent(r)
	if err != nil {
		raiseError(w, err)
		return
	}

	if params.Type == feature.URLVerification {
		mSlack.VerificateUrl(w, *params)
		return
	}

	route, err := routes.GetRoute(params.RequestKey)
	if err != nil {
		raiseError(w, err)
		return
	}

	route.Feature.RunEvent(*params)
}
