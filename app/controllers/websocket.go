package controllers

import (
	"github.com/revel/revel"
	"golang.org/x/net/websocket"
	"fmt"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) RoomSocket(ws *websocket.Conn) revel.Result {
	if ws == null {
		return nil
	}

	subscription := Subscribe()
	defer subscription.Cancel()
}

