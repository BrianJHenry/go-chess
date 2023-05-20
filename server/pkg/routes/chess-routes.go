package routes

import (
	"github.com/BrianJHenry/go-chess/server/pkg/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func BuildRoutes(app *fiber.App) {

	// non websocket routes

	// websocket
	manager := controllers.NewManager()

	app.Use("/ws", manager.UseWebsocket)

	app.Get("/ws/:id", websocket.New(manager.ServeWebsocket))

}
