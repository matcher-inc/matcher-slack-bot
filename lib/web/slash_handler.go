package web

import (
	"go-bot-test/config/routes"
	mSlack "go-bot-test/lib/m_slack"
	"net/http"
)

func handleSlash(w http.ResponseWriter, r *http.Request) {
	params, err := mSlack.ParseSlash(r)
	if err != nil {
		raiseError(w, err)
		return
	}

	route, err := routes.GetRoute(params.RequestKey)
	if err != nil {
		raiseError(w, err)
		return
	}

	route.Feature.RunEvent(*params)
}
