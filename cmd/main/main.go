package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/BrianJHenry/go-chess/pkg/models"
)

func main() {
	start := time.Now()
	game := models.NewChessGame()
	for i := 0; i < 40; i++ {
		possibleMoves := game.CurrentState.EnumerateMoves()
		game.ExecuteMove(possibleMoves[rand.Intn(len(possibleMoves))])
	}
	fmt.Printf("Completed %v half moves in %v\n", 100, time.Since(start))
}
