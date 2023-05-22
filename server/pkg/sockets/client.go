package sockets

import (
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

	var err error
	var move = &moveToSend{}
	for {
		err = c.Conn.ReadJSON(move)
		if err != nil {
			log.Println(err)
			return
		}
		c.Game.RecieveMove <- *move
	}
}
