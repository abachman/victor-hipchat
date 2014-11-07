// Based on the Slack library for victor https://github.com/brettbuddin/victor

package hipchatRobot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/abachman/hipchat-go/hipchat"
	"github.com/brettbuddin/victor/pkg/chat"
)

func init() {
	chat.Register("hipchat", func(r chat.Robot) chat.Adapter {
		team := os.Getenv("VICTOR_HIPCHAT_ROOMS")
		token := os.Getenv("VICTOR_HIPCHAT_TOKEN")

		if team == "" || token == "" {
			log.Println("The following environment variables are required:")
			log.Println("VICTOR_HIPCHAT_ROOMS, VICTOR_HIPCHAT_TOKEN")
			os.Exit(1)
		}

		client := hipchat.NewClient(token)

		return &hipchatRobot{
			robot:  r,
			client: client,
		}
	})
}

type hipchatRobot struct {
	robot  chat.Robot
	client *hipchat.Client
}

func (h *hipchatRobot) Run() {
	h.robot.HTTP().HandleFunc("/patbot/hipchat-webhook", func(w http.ResponseWriter, r *http.Request) {
		// debug message JSON
		body, _ := ioutil.ReadAll(r.Body)

		msg := &WebhookMessage{}
		err := json.NewDecoder(strings.NewReader(string(body))).Decode(msg)
		if err != nil {
			fmt.Println("[hipchat-webhook] failed to decode message:", err)
			fmt.Println(string(body))
		}

		// now communicate message to victor
		h.robot.Receive(&message{
			userID:      strconv.Itoa(msg.Item.Message.From.ID),
			userName:    msg.Item.Message.From.Name,
			channelID:   strconv.Itoa(msg.Item.Room.ID),
			channelName: msg.Item.Room.Name,
			text:        msg.Item.Message.Message,
		})
	}).Methods("POST")
}

func (s *hipchatRobot) Send(channelID, msg string) {
	//	Color         string
	//	Message       string
	//	Notify        bool
	//	MessageFormat string
	resp, err := s.client.Room.Notification(channelID, &hipchat.NotificationRequest{
		Message:       msg,
		Notify:        true,
		MessageFormat: "html",
	})

	if err != nil {
		log.Println("error sending to chat:", err, resp.Body)
	}
}

func (s *hipchatRobot) Stop() {
	// no need, webhook listeners are stopped with the HTTP server
}

// victor boilerplate
type message struct {
	userID, userName, channelID, channelName, text string
}

func (m *message) UserID() string {
	return m.userID
}

func (m *message) UserName() string {
	return m.userName
}

func (m *message) ChannelID() string {
	return m.channelID
}

func (m *message) ChannelName() string {
	return m.channelName
}

func (m *message) Text() string {
	return m.text
}
