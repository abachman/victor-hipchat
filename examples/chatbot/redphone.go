package main

import (
	"fmt"
	"net/http"
	"os"
)

// e.g. r13402
var (
	redphoneUrl      = os.Getenv("BOT_REDPHONE_URL")
	redphoneEndpoint = "http://redphonesupport.dev/r/"
)

//// Private API

// Ask Redphone to send a pat.f53.co/say message with a link to the given Redphone Ticket
func RequestTicketMessage(roomID string, identifier string) string {
	ticket_link_url := fmt.Sprintf("%v/patbot/ticket_link?r=%v&roomID=%v", redphoneUrl, identifier, roomID)

	resp, err := http.Get(ticket_link_url)
	if err != nil {
		m := fmt.Sprintf("Redphone Ticket %s: %s", identifier, redphoneEndpoint+identifier)
		return fmt.Sprintf("%v — Couldn't reach Redphone with error: %v", m, err)
	}
	defer resp.Body.Close()

	return ""
}
