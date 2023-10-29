package main

import (
	"log"
	"net/http"

	"github.com/chanon-mike/stupid-discord-bot/discord"
)

func main() {
	go startKeepAliveServer()
	discord.Start()
}

// Healthcheck server for 5-min interval in order to prevent Render cold start
func startKeepAliveServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Alive"))
	})

	log.Println("Start keeping alive server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error, could not start server: %v\n", err)
	}
}
