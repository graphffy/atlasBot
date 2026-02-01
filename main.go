package main

import (
	"atlasBot/server"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", server.Handler)

	err := http.ListenAndServe(":9003", nil)

	if err != nil {
		log.Fatal(err)
	}
}
