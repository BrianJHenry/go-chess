package sockets

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Game *Game
}

func (c *Client) Read() {
	defer func() {
		c.Conn.Close()
		c.Game.Unregister <- c
	}()

	var err error
	var move = &APIMove{}
	log.Println("Begin read")
	for {
		err = c.Conn.ReadJSON(move)
		if err != nil {
			log.Println(err)
			return
		}
		c.Game.RecieveMove <- *move
	}
}
