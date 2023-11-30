package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Message represents a chat message.
type Message struct {
	Username string `json:"username"`
	Room     string `json:"room"`
	Text     string `json:"text"`
}

// Client represents a connected user.
type Client struct {
	Username string
	Room     string
	Conn     *websocket.Conn
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients   = make(map[*websocket.Conn]Client)
	clientsMu sync.Mutex
)

func wsHandler(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), c.Response().Header())

	if err != nil {
		return err
	}

	username := c.QueryParam("username")
	room := c.QueryParam("room")

	client := Client{
		Username: username,
		Room:     room,
		Conn:     conn,
	}

	clientsMu.Lock()
	clients[conn] = client
	clientsMu.Unlock()

	log.Printf("User %s joined room %s", username, room)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		// Broadcast the message to all clients in the same room
		clientsMu.Lock()
		for _, c := range clients {
			if c.Room == msg.Room {
				err := c.Conn.WriteJSON(msg)
				if err != nil {
					log.Printf("Error writing message: %v", err)
					break
				}
			}
		}
		clientsMu.Unlock()
	}

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/ws", wsHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
