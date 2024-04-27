package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexmgriffiths/habbo-go/game"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/alexmgriffiths/habbo-go/utils"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var databaseManager *managers.DatabaseManager = managers.NewDatabaseManager()
var gameManager *managers.GameManager = managers.NewGameManager()
var logger *utils.Logger = utils.NewLogger()

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

type IncomingPacket struct {
	header uint16
	length int
	buffer utils.ByteBuf
	data   []byte
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
		var packet IncomingPacket = IncomingPacket{header, size, *buffer, message}

		logger.Success("[INCOMING] %d", packet.header)

		packet.buffer.ReadInt()   // Not sure what this one is
		packet.buffer.ReadShort() // Header again

		if packet.header == 2419 {

			ticket := packet.buffer.ReadString()
			logger.Success("TICKET: %s", ticket)

			response := utils.NewByteBuf([]byte{})
			response.WriteShort(2491)

			habbo, err := game.NewHabbo(databaseManager.GetConnection(), c, ticket)
			gameManager.AddClient(c, habbo)

			var testConnection *game.Habbo = gameManager.GetClient(c)

			if err != nil {
				log.Fatal("Error", err)
			}

			log.Println(habbo)

			testConnection.GetConnection().WriteMessage(websocket.BinaryMessage, response.Wrap())
		} else if packet.header == 357 {

			currentHabbo := gameManager.GetClient(c)
			response := utils.NewByteBuf([]byte{})
			response.WriteShort(2725)

			response.WriteInt(currentHabbo.GetId())
			response.WriteString(currentHabbo.GetUsername())
			response.WriteString(currentHabbo.GetLook())
			response.WriteString(currentHabbo.GetGender())
			response.WriteString(currentHabbo.GetMotto())
			response.WriteString(currentHabbo.GetUsername())
			response.WriteBool(false)
			response.WriteInt(0)
			response.WriteInt(0)
			response.WriteInt(0)
			response.WriteBool(false)
			response.WriteString("01-01-1970 00:00:00")
			response.WriteBool(false)
			response.WriteBool(false)

			c.WriteMessage(websocket.BinaryMessage, response.Wrap())

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
