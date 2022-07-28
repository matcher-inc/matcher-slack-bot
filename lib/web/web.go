package web

import (
	"log"
	"net/http"
)

func RunServer() {
	handleRequest()
	serve()
}

func handleRequest() {
	http.HandleFunc("/slack/events", handleEvent)
	http.HandleFunc("/slack/slashes", handleSlash)
	http.HandleFunc("/slack/actions", handleAction)
}

func serve() {
	log.Println("[INFO] Server listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func raiseError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
}
