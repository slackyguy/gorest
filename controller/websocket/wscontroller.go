package websocket

import (
	"net/http"

	"github.com/slackyguy/gorest/routing"
	"github.com/slackyguy/gorest/websocket"
)

func init() {
	hub := websocket.NewHub()
	go hub.Run()

	routing.HTTP.RegisterHandler("/ws", func(
		response http.ResponseWriter, request *http.Request) {
		websocket.ServeWs(hub, response, request)
	})
}

// DoNothing (NOP)
func DoNothing() {}
