package main

import (
	"log"

	"github.com/BrianJHenry/go-chess/server/pkg/chesssockets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func setupRoutes(app *fiber.App) {

	games := make(map[string]*chesssockets.Game)

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
				client := &chesssockets.Client{
					Conn: conn,
					Game: game,
				}
				game.Register <- client
				client.Read()
			} else {
				// Respond with Game full
			}
		} else {
			// Respond with invalid ID
		}
	}))
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
