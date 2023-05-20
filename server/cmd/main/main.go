package main

import (
	"log"

	"github.com/BrianJHenry/go-chess/server/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.BuildRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
