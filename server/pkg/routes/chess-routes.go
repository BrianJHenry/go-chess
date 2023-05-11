package routes

import (
	"github.com/BrianJHenry/go-chess/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func BuildRoutes(router *gin.Engine) {
	router.GET("/play", controllers.GetGame)
	router.GET("/play/:id/updateState", controllers.UpdateState)
	router.PUT("/play/:id/playMove", controllers.PlayMove)
}
