package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		From string `json:"from"`
		To string 	`json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestedPayload mailMessage

	err := app.readJSON(w, r, &requestedPayload)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	msg := Message {
		From: requestedPayload.From,
		To: requestedPayload.To,
		Subject: requestedPayload.Subject,
		Data: requestedPayload.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	if err != nil {
		log.Println(err)
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse {
		Error: false,
		Message: "sent to " + requestedPayload.To,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}