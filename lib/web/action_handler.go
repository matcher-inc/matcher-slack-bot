package web

import (
	"go-bot-test/config/routes"
	mSlack "go-bot-test/lib/m_slack"
	"net/http"
)

func handleAction(w http.ResponseWriter, r *http.Request) {
	params, err := mSlack.ParseAction(r)
	if err != nil {
		raiseError(w, err)
		return
	}

	route, err := routes.GetRoute(params.FeaturePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	route.Feature.RunAction(*params)
}
