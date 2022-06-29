package web

import (
	"log"
	"net/http"
)

func RunServer() {
	handleRequest()

	log.Println("[INFO] Server listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleRequest() {
	http.HandleFunc("/slack/events", handleEvent)
	http.HandleFunc("/slack/actions", handleAction)
}
