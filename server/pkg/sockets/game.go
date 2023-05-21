package sockets

import (
	"fmt"

	"github.com/BrianJHenry/go-chess/server/pkg/models"
)

type Game struct {
	// game info
	NumberOfPlayers int

	Clients []*Client

	// websocket handling
	Register    chan *Client
	Unregister  chan *Client
	RecieveMove chan Message
}

func NewGame(numberOfPlayers int) *Game {
	return &Game{
		NumberOfPlayers: numberOfPlayers,
		Clients:         make([]*Client, numberOfPlayers),
		Register:        make(chan *Client),
		Unregister:      make(chan *Client),
		RecieveMove:     make(chan Message),
	}
}

func (game *Game) Start() {

	// current number of clients to check if we have enough to start our game
	currentClientCount := 0

	// create new game
	chessGame := models.NewChessGame()
	gameOver := false

	for !gameOver {
		select {
		case client := <-game.Register:
			game.Clients[currentClientCount] = client
			currentClientCount += 1
			fmt.Println("Number of players in lobby: ", len(game.Clients))
			// if we have enough players start the game
			if currentClientCount == game.NumberOfPlayers {
				// send over board to all
				// send whose turn to all
				for i, c := range game.Clients {
					fmt.Println(i, c)
					// TODO: send board state, turn, and previous moves as json to client
				}
			}
		case client := <-game.Unregister:
			for i, c := range game.Clients {
				if c == client {
					game.Clients[i] = nil
				} else {
					// set winner to whichever client did not disconnect
					winner := ""
					if i == 0 {
						winner = "White"
					} else {
						winner = "Black"
					}
					client.Conn.WriteMessage(1, []byte(fmt.Sprintf("Opponent Disconnected. %s wins!", winner)))
				}
			}
			gameOver = true
		case move := <-game.RecieveMove:
			// TODO: check that move was in possible moves

			// TODO: execute move

			// TODO: send back updated state to all players

			for _, client := range game.Clients {
				client.Conn.WriteMessage(move.Type, []byte(move.Body))
			}
		}

		// TODO: check if the previous move ended the game

		// TODO: if so, set gameOver = true
		// TODO: if so, send end game message to players

		// TODO: check if the computer needs to make a move

		// TODO: enumerate moves and make move for computer
		// TODO: check if that ends game
		// TODO: if so send endd game message to players

		// TODO: if the computer doesn't need to make a move

		// TODO: then enumerate moves for the color whose turn it is
		// TODO: send possible moves to appropriate client
	}
}
