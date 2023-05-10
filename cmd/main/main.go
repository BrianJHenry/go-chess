package main

import (
	"fmt"
	"time"

	"github.com/BrianJHenry/go-chess/pkg/models"
)

func main() {
	chessState := models.StartingState()
	start := time.Now()
	moves := chessState.EnumerateMoves()
	duration := time.Since(start)
	for _, val := range moves {
		fmt.Printf("%b\n", val)
	}
	fmt.Printf("Number of moves found: %v\n", len(moves))
	fmt.Printf("Execution time of enumerate moves: %v\n", duration)
}
