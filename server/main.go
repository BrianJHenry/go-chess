package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/BrianJHenry/go-chess/server/pkg/sockets"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var games = make(map[string]*sockets.Game)

func UnregisterGame(gameID string) {
	delete(games, gameID)
}

func setupRoutes(app *fiber.App) {

	app.Get("/findGame/:numPlayers", func(c *fiber.Ctx) error {
		numberOfPlayers, err := strconv.Atoi(c.Params("numPlayers"))
		if err != nil {
			log.Println("Invalid Find Game Request")
			return c.Status(404).SendString(err.Error())
		}
		if numberOfPlayers == 2 {
			for key, element := range games {
				if element.NumberOfPlayers == 2 && len(element.Clients) < 2 {
					return c.SendString(key)
				}
			}
		}
		var randomKey string = fmt.Sprint(rand.Intn(10000000))
		for _, ok := games[randomKey]; ok; {
			randomKey = fmt.Sprint(rand.Intn(10000000))
		}
		newGame := sockets.NewGame(numberOfPlayers, randomKey, func(gameID string) {
			delete(games, gameID)
			log.Println("Deleting game")
		})
		games[randomKey] = newGame
		return c.SendString(randomKey)
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
