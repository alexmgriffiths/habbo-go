package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/communication/incoming/handshake"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/alexmgriffiths/habbo-go/utils"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var databaseManager *managers.DatabaseManager = managers.NewDatabaseManager()
var gameManager *managers.GameManager = managers.NewGameManager(databaseManager.GetConnection())
var logger *utils.Logger = utils.NewLogger()

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

type PacketHandler interface {
	Handle(gm *managers.GameManager, incomingPacket *incoming.IncomingPacket, c *websocket.Conn) error
}

var packetHandlers = map[uint16]PacketHandler{
	4000: &handshake.ClientHelloEvent{},
	2419: &handshake.SSOTicketEvent{},
	357:  &handshake.GetUserInfoEvent{},
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsh.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error %s when upgrading connection to websocket", err)
		return
	}

	defer c.Close() // Runs after function completes
	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Println("Connection closed by client")
				return
			}
			log.Printf("Error reading message: %s", err)
			return
		}

		if messageType == websocket.TextMessage {
			log.Printf("Failed to read message because it was type text")
			return
		}

		header := binary.BigEndian.Uint16((message[4:6]))
		size := len(message)
		buffer := utils.NewByteBuf(message)
		packet := incoming.NewIncomingPacket(header, size, *buffer, message)

		logger.Success("[INCOMING] %d", packet.GetHeader())

		packet.GetBuffer().ReadInt()   // Not sure what this one is
		packet.GetBuffer().ReadShort() // Header again

		handler, ok := packetHandlers[header]
		if ok {
			handler.Handle(gameManager, packet, c)
		} else {
			logger.Error("Unhandled packet header: %d", header)
		}

	}

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}
	http.Handle("/", webSocketHandler)
	logger.Success("Connected!")
	log.Print("Starting server...")

	var emuHost string = os.Getenv("emu_host")
	var emuPort string = os.Getenv("emu_port")

	var emuStr string = fmt.Sprintf("%s:%s", emuHost, emuPort)
	log.Println(emuStr)

	log.Fatal(http.ListenAndServe(emuStr, nil))
}
