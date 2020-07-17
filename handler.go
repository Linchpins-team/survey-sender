package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	Config
	sender Sender
}

func NewHandler() Handler {
	return Handler{
		Config: Configure,
		sender: Sender{Configure},
	}
}

func (h Handler) newSurvey(w http.ResponseWriter, r *http.Request) {
	var survey Survey
	dec := json.NewDecoder(r.Body)

	var err error
	if err = dec.Decode(&survey); err != nil {
		http.Error(w, err.Error(), 403)
		return
	}
	survey.CreatedAt = time.Now()
	if !survey.CheckAvailable() {
		http.Error(w, "This survey is not available", 403)
		return
	}
	if err = Save(survey); err != nil {
		panic(err)
	}
	if err = h.sender.Send(survey); err != nil {
		panic(err)
	}
}
