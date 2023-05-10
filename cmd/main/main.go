package main

import (
	"fmt"

	"github.com/BrianJHenry/go-chess/pkg/models"
)

func main() {
	chessState := models.StartingState()
	moves := chessState.EnumerateMoves()
	for _, val := range moves {
		fmt.Printf("%b\n", val)
	}
}
