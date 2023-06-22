package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/BrianJHenry/go-chess/server/pkg/sockets"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
		if numberOfPlayers != 1 && numberOfPlayers != 2 {
			log.Println("Invalid number of players")
			return c.Status(404).SendString("Invalid number of players.")
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
		log.Println("Creating new game.")
		games[randomKey] = newGame
		go newGame.Start()
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
		log.Println("Requested websocket with ID: ", id)

		if game, ok := games[id]; ok {
			log.Println("Number of clients currently in game: ", len(game.Clients))
			if len(game.Clients) < game.NumberOfPlayers {
				client := &sockets.Client{
					Conn: conn,
					Game: game,
				}
				log.Print("About to register...")
				game.Register <- client
				log.Println("Registering client and starting read.")
				client.Read()
			} else {
				log.Printf("Game with ID: %v is full.", id)
				conn.Conn.WriteMessage(1, []byte(fmt.Sprintf("Game with ID: %v is full.", id)))
			}
		} else {
			log.Println("Invalid game ID.")
			conn.Conn.WriteMessage(1, []byte("Invalid game ID."))
		}
	}))
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	setupRoutes(app)

	log.Println("Starting app.")
	log.Fatal(app.Listen(":3000"))
}
