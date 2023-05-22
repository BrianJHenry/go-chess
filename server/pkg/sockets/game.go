package sockets

import (
	"fmt"

	"github.com/BrianJHenry/go-chess/server/pkg/models"
)

// setup move information to be sent across websockets
var moveTypesArray [8]string = [8]string{"N", "S", "L", "P", "Q", "R", "B", "K"}
var moveTypesMap map[string]int = map[string]int{
	"N": 0,
	"S": 1,
	"L": 2,
	"P": 3,
	"Q": 4,
	"R": 5,
	"B": 6,
	"K": 7,
}

type moveToSend struct {
	moveType  string `json:"moveType"`
	oldSquare int    `json:"oldSquare"`
	newSquare int    `json:"newSquare"`
}

func convertToMoveToSend(move models.Move) moveToSend {
	moveType := moveTypesArray[move.Type]
	oldSquare := move.OldSquare.Row*8 + move.OldSquare.Col
	newSquare := move.NewSquare.Row*8 + move.NewSquare.Col

	return moveToSend{
		moveType:  moveType,
		oldSquare: oldSquare,
		newSquare: newSquare,
	}
}

func convertToMove(move moveToSend) models.Move {
	moveType := moveTypesMap[move.moveType]
	oldSquare := models.Location{
		Row: move.oldSquare / 8,
		Col: move.oldSquare % 8,
	}
	newSquare := models.Location{
		Row: move.newSquare / 8,
		Col: move.newSquare % 8,
	}

	return models.Move{
		Type:      models.MoveType(moveType),
		OldSquare: oldSquare,
		NewSquare: newSquare,
	}
}

// setup states to be sent across websockets
type stateToSend struct {
	turn          bool
	board         []int8
	previousMoves []moveToSend
}

func convertToStateToSend(game models.ChessGame, ownColor int) stateToSend {
	var board = make([]int8, 0, 64)
	for _, row := range game.CurrentState.Board {
		sliceRow := row[:]
		board = append(board, sliceRow...)
	}
	var turn bool
	if ownColor == int(game.CurrentState.Turn) {
		turn = true
	} else {
		turn = false
	}
	var previousMoves = make([]moveToSend, 0, len(game.Moves))
	for _, move := range game.Moves {
		convertedMove := convertToMoveToSend(move)
		previousMoves = append(previousMoves, convertedMove)
	}

	return stateToSend{
		turn:          turn,
		board:         board,
		previousMoves: previousMoves,
	}
}

// actual game logic
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
				for i, c := range game.Clients {
					fmt.Println(i, c)
					startingState := convertToStateToSend(chessGame, i)
					client.Conn.WriteJSON(startingState)
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
