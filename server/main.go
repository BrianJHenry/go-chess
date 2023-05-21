package main

import (
	"fmt"
	"log"

	"github.com/BrianJHenry/go-chess/server/pkg/sockets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func setupRoutes(app *fiber.App) {

	games := make(map[string]*sockets.Game)

	app.Get("/findGame", func(c *fiber.Ctx) error {
		// TODO: implement searching for free game
		// TODO: if no game found, implement adding game to games and starting it
		return nil
	})

	app.Use("/game", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/game/:id", websocket.New(func(conn *websocket.Conn) {
		id := conn.Params("id")

		if game, ok := games[id]; ok {
			if len(game.Clients) < game.NumberOfPlayers {
				client := &sockets.Client{
					Conn: conn,
					Game: game,
				}
				game.Register <- client
				client.Read()
			} else {
				conn.Conn.WriteMessage(1, []byte(fmt.Sprintf("Game with ID: %v is full.", id)))
			}
		} else {
			conn.Conn.WriteMessage(1, []byte("Invalid game ID."))
		}
	}))
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
