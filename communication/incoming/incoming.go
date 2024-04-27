package incoming

import (
	"github.com/gorilla/websocket"
)

type PacketHandler interface {
	Handle(c *websocket.Conn) error
}
