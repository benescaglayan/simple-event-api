package main

import (
	"event-api/repository"
	"net/http"
)

func main() {
	repository.EventRepository.Initialize()

	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}
