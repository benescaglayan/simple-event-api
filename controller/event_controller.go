package controller

import (
	"encoding/json"
	"event-api/model"
	"event-api/service"
	"net/http"
)

func (ctr *Controller) ProcessEvent(w http.ResponseWriter, r *http.Request) {
	var eventRequest model.EventRequest

	err := json.NewDecoder(r.Body).Decode(&eventRequest)
	if err != nil {
		configureHttpResponse(w, "Invalid JSON object received.", http.StatusBadRequest)
		return
	}

	err = service.EventProcessor.ProcessEvent(eventRequest)
	if err != nil {
		configureHttpResponse(w, err.Error(), http.StatusBadRequest)
	} else {
		configureHttpResponse(w, "Event processed successfully.", http.StatusOK)
	}
}

func configureHttpResponse(w http.ResponseWriter, message string, statusCode int) {
	var eventResponse model.EventResponse
	eventResponse.Message = message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(eventResponse)
}
