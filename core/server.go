package core

import (
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/communication/incoming/handshake"
	"github.com/alexmgriffiths/habbo-go/communication/incoming/navigator"
	"github.com/alexmgriffiths/habbo-go/communication/incoming/rooms"
	incomingRoomUser "github.com/alexmgriffiths/habbo-go/communication/incoming/rooms/users"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/alexmgriffiths/habbo-go/utils"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var logger *utils.Logger = utils.NewLogger()
var managersHolder *managers.Managers

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

type PacketHandler interface {
	Handle(managers *managers.Managers, incomingPacket *incoming.IncomingPacket, c *websocket.Conn) error
}

// Ampersand because each of the handlers take a pointer for the PacketHandler:
/*
	func (e *ClientHelloEvent) Handle(...

	The key part to notice is the context of (e *ClientHelloEvent) that's why we need pointers
	We are using a pointer since we don't need a copy of the events, we can just reuse the exact same one over and over
*/
var packetHandlers = map[uint16]PacketHandler{
	4000: &handshake.ClientHelloEvent{},
	2419: &handshake.SSOTicketEvent{},
	357:  &handshake.GetUserInfoEvent{},
	2110: &navigator.NavigatorInitEvent{},
	249:  &navigator.NavigatorSearchEvent{},
	2312: &rooms.OpenFlatConnection{},
	3898: &rooms.GetRoomDataEvent{}, // Called when entering room from hotel view
	2300: &rooms.GetRoomDataEvent{}, // Called when moving from one room to another
	3320: &incomingRoomUser.MoveAvatarEvent{},
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
			handler.Handle(managersHolder, packet, c)
		} else {
			logger.Error("Unhandled packet header: %d", header)
		}

	}

}

func Start() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}
	http.Handle("/", webSocketHandler)

	// Register managers
	databaseManager := managers.NewDatabaseManager()
	gameManager := managers.NewGameManager(databaseManager.GetConnection())
	roomsManager := managers.NewRoomManager(databaseManager.GetConnection(), logger)

	managersHolder = managers.NewManagers(databaseManager.GetConnection(), roomsManager, gameManager)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		managersHolder.GetRoomManager().LoadCategories()
	}()

	go func() {
		defer waitGroup.Done()
		managersHolder.GetRoomManager().LoadLayouts()
		managersHolder.GetRoomManager().LoadCustomLayouts()
		managersHolder.GetRoomManager().LoadRooms()
	}()

	waitGroup.Wait()

	go func() {
		for {
			// Call StartCycles every 500 milliseconds
			managersHolder.GetGameManager().StartCycles()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	log.Print("Starting server...")

	var emuHost string = os.Getenv("emu_host")
	var emuPort string = os.Getenv("emu_port")

	var emuStr string = fmt.Sprintf("%s:%s", emuHost, emuPort)
	log.Println(emuStr)

	log.Fatal(http.ListenAndServe(emuStr, nil))
}

func GetManages() *managers.Managers {
	return managersHolder
}
