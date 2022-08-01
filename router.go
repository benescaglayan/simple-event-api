package main

import (
	"event-api/controller"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	c := controller.NewController()

	switch r.Method {
	case "POST":
		c.ProcessEvent(w, r)
	}
}
