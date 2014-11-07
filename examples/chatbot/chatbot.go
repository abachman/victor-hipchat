package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/abachman/victor-hipchat/hipchat"
	"github.com/brettbuddin/victor"
)

func main() {
	bot := victor.New(victor.Config{
		Name:         "pat",
		ChatAdapter:  "hipchat",
		StoreAdapter: "memory",
		HTTPAddr:     ":9900",
	})

	bot.HTTP().HandleFunc("/say", func(w http.ResponseWriter, r *http.Request) {
		message := r.FormValue("message")
		roomID := r.FormValue("room")
		bot.Chat().Send(roomID, message)
	})

	// presence check
	bot.HandleFunc("pat\\?", func(s victor.State) {
		s.Chat().Send(s.Message().ChannelID(), fmt.Sprintf("I'm here, %s", s.Message().UserName()))
	})

	// response check
	bot.HandleCommandFunc("hello|hi|howdy", func(s victor.State) {
		s.Chat().Send(s.Message().ChannelID(), fmt.Sprintf("Hello yourself, %s", s.Message().UserName()))
	})

	bot.HandleFunc(`(?:^r| r)([0-9]{4,6})([^0-9]|$)`, func(s victor.State) {
		reply := RequestTicketMessage(s.Message().ChannelID(), s.Params()[0])
		if len(reply) > 0 {
			bot.Chat().Send(s.Message().ChannelID(), reply)
		}
	})

	go bot.Run()

	// clean shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs

	bot.Stop()
}
