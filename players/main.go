package main

import (
	"go_intro/players/server"
	"go_intro/players/store"
	"log"
	"net/http"
)

func main() {
	store := store.New()
	server := server.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
