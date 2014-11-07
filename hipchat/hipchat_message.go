package hipchatRobot

type User struct {
	ID          int    `json:"id"`
	MentionName string `json:"mention_name"`
	Name        string `json:"name"`
}

type Message struct {
	ID      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Message string `json:"message,omitempty"`
	From    User   `json:"from,omitempty"`
}

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MessageData struct {
	Message Message `json:"message"`
	Room    Room    `json:"room"`
}

// top level hipchat webhook JSON object
type WebhookMessage struct {
	Event         string      `json:"string"`
	Item          MessageData `json:"item"`
	WebhookID     int         `json:"webhook_id"`
	OauthClientID string      `json:"oauth_client_id"`
}
