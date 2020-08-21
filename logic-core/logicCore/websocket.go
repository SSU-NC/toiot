package logicCore
/*
import (
	"fmt"

	"github.com/gorilla/websocket"
)

type WebSocket struct {
	conn *websocket.Conn
}

func NewWebSocket(c *websocket.Conn) *WebSocket{
	return &WebSocket{
		conn : c,
	}
	 
	go func() {
		ch
	}
}

func (w *WebSocket) send(message string) {
	var msg string
	msg = `{alarm:` + message + `}`
	if err := w.conn.WriteJSON(msg); err != nil {
			fmt.Println("alarm websocket error")
	}
}*/