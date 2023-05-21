package sockets

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Conn *websocket.Conn
	Game *Game
}

type Message struct {
	Type int
	Body string
}

func (c *Client) Read() {
	defer func() {
		c.Conn.Close()
		c.Game.Unregister <- c
	}()

	var (
		messageType int
		message     []byte
		err         error
	)
	for {
		messageType, message, err = c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		move := Message{Type: messageType, Body: string(message)}
		c.Game.RecieveMove <- move
		fmt.Printf("Message Recieved: %+v", move)
	}
}
