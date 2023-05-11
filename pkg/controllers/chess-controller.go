package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGame(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Get Game Endpoint"})
}

func UpdateState(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Update State Endpoint"})
}

func PlayMove(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Play Move Endpoint"})
}
