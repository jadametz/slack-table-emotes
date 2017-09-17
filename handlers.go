package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// SlackAttachment represents an optional piece of the SlackResponse
type SlackAttachment struct {
	AttachmentText string `json:"text"`
}

// SlackPayload represents an incoming slash command payload
type SlackPayload struct {
	Token       string `schema:"token"`
	TeamID      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelID   string `schema:"channel_id"`
	ChannelName string `schema:"channel_name"`
	UserID      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Command     string `schema:"command"`
	Text        string `schema:"text"`
	ResponseURL string `schema:"response_url"`
}

// SlackResponse is the payload returned to Slack
type SlackResponse struct {
	ResponseType string            `json:"response_type"`
	Text         string            `json:"text"`
	Attachments  []SlackAttachment `json:"attachments"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("/healthz")
	return
}

func tableHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "Unable to parse Slack form payload", http.StatusBadRequest)
		return
	}

	slackData := new(SlackPayload)

	if err := decoder.Decode(slackData, r.Form); err != nil {
		log.Println(err)
		http.Error(w, "Unable to decode Slack form payload", http.StatusInternalServerError)
		return
	}

	log.Printf("Team: %s User: %s Command: %s", slackData.TeamID, slackData.UserID, slackData.Text)

	switch {
	case slackData.Text == "flip":
		tableResponder(w, tableFlip)
	case slackData.Text == "catch":
		tableResponder(w, tableCatch)
	default:
		w.Write([]byte("Did you want me to `flip` or `catch` the table?"))
	}
}

func tableResponder(w http.ResponseWriter, action TableAction) {
	attachments := []SlackAttachment{}
	if os.Getenv("ATTACHMENTS") == "yes" {
		attachments = append(attachments, SlackAttachment{AttachmentText: action.Description})
	}

	payload, err := json.Marshal(SlackResponse{
		ResponseType: "in_channel",
		Text:         action.Emote,
		Attachments:  attachments,
	})
	if err != nil {
		log.Println("Unalbe to marhsal JSON response")
		http.Error(w, "Unable to marshal JSON response", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
