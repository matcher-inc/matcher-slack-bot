package main

import (
	"go-bot-test/app/actions"
	"go-bot-test/app/constants"
	"go-bot-test/app/events"
	"go-bot-test/lib/listner"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/slack/events", listner.ListenEvent([]listner.EventRoute{
		listner.MentionEventRoute{Command: constants.PingEvent, Handler: events.PingEvent},
		listner.MentionEventRoute{Command: constants.DeployEvent, Handler: events.DeployEvent},
	}))

	http.HandleFunc("/slack/actions", listner.ListenAction([]listner.ActionRoute{
		listner.InteractionTypeBlockActionRoute{ActionId: constants.SelectVersionAction, Handler: actions.SelectVersionAction},
		listner.InteractionTypeBlockActionRoute{ActionId: constants.ConfirmDeploymentAction, Handler: actions.ConfirmDeploymentAction},
	}))

	log.Println("[INFO] Server listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
