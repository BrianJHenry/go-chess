package main

import (
	"github.com/BrianJHenry/go-chess/server/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.BuildRoutes(router)

	router.Run(":8080")
}
