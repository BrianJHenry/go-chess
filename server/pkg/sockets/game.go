package sockets

import (
	"fmt"
	"math/rand"

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
	MoveType  string `json:"moveType"`
	OldSquare int    `json:"oldSquare"`
	NewSquare int    `json:"newSquare"`
}

func convertToMoveToSend(move models.Move) moveToSend {
	moveType := moveTypesArray[move.Type]
	oldSquare := move.OldSquare.Row*8 + move.OldSquare.Col
	newSquare := move.NewSquare.Row*8 + move.NewSquare.Col

	return moveToSend{
		MoveType:  moveType,
		OldSquare: oldSquare,
		NewSquare: newSquare,
	}
}

func convertToMove(move moveToSend) models.Move {
	moveType := moveTypesMap[move.MoveType]
	oldSquare := models.Location{
		Row: move.OldSquare / 8,
		Col: move.OldSquare % 8,
	}
	newSquare := models.Location{
		Row: move.NewSquare / 8,
		Col: move.NewSquare % 8,
	}

	return models.Move{
		Type:      models.MoveType(moveType),
		OldSquare: oldSquare,
		NewSquare: newSquare,
	}
}

// setup states to be sent across websockets
type stateToSend struct {
	Turn          bool         `json:"turn"`
	Board         []int8       `json:"board"`
	PreviousMoves []moveToSend `json:"previousMoves"`
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
	var previousMoves = make([]moveToSend, 0, len(game.MoveHistory))
	for _, move := range game.MoveHistory {
		convertedMove := convertToMoveToSend(move)
		previousMoves = append(previousMoves, convertedMove)
	}

	return stateToSend{
		Turn:          turn,
		Board:         board,
		PreviousMoves: previousMoves,
	}
}

// actual game logic
type Game struct {
	GameID string
	Delete func(id string)

	// game info
	NumberOfPlayers int

	Clients []*Client

	// websocket handling
	Register    chan *Client
	Unregister  chan *Client
	RecieveMove chan moveToSend
}

func NewGame(numberOfPlayers int, gameID string, delete func(id string)) *Game {
	return &Game{
		GameID:          gameID,
		Delete:          delete,
		NumberOfPlayers: numberOfPlayers,
		Clients:         make([]*Client, numberOfPlayers),
		Register:        make(chan *Client),
		Unregister:      make(chan *Client),
		RecieveMove:     make(chan moveToSend),
	}
}

func (game *Game) Start() {
	defer func() {
		game.Delete(game.GameID)
	}()

	// create new game
	chessGame := models.NewChessGame()
	gameOver := false

	for !gameOver {
		select {
		case client := <-game.Register:
			game.Clients = append(game.Clients, client)
			fmt.Println("Number of players in lobby: ", len(game.Clients))
			// if we have enough players start the game
			if len(game.Clients) == game.NumberOfPlayers {
				for i, c := range game.Clients {
					fmt.Println(i, c)
					startingState := convertToStateToSend(chessGame, i)
					client.Conn.WriteJSON(startingState)
					if i == 0 {
						// send over possible moves to white
						client.Conn.WriteJSON(chessGame.PossibleMoves)
					}
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
			// check that move was in possible moves
			tryMove := convertToMove(move)
			isAllowedMove := false
			for _, possibleMove := range chessGame.PossibleMoves {
				if tryMove == possibleMove {
					isAllowedMove = true
					break
				}
			}
			if !isAllowedMove {
				game.Clients[chessGame.CurrentState.Turn].Conn.WriteMessage(1, []byte("Invalid Move Attempted."))
				break
			}

			// execute move
			chessGame.ExecuteMoveOnGame(tryMove)

			// send back updated state
			for i, client := range game.Clients {
				client.Conn.WriteJSON(convertToStateToSend(chessGame, i))
			}

			// check if game is ended

			// TODO: Refactor logic
			if chessGame.Winner == models.Stalemate {
				for _, client := range game.Clients {
					client.Conn.WriteMessage(1, []byte("Stalemate"))
				}
				gameOver = true
				break
			} else if chessGame.Winner == models.WhiteWins {
				for _, client := range game.Clients {
					client.Conn.WriteMessage(1, []byte("White Wins!"))
				}
				gameOver = true
				break
			} else if chessGame.Winner == models.BlackWins {
				for _, client := range game.Clients {
					client.Conn.WriteMessage(1, []byte("Black Wins!"))
				}
				gameOver = true
				break
			}

			// game continues... (pass move to player or execute computer move)
			if game.NumberOfPlayers == 1 {
				// execute random move
				chessGame.ExecuteMoveOnGame(chessGame.PossibleMoves[rand.Intn(len(chessGame.PossibleMoves))])

				// send back updated state
				for i, client := range game.Clients {
					client.Conn.WriteJSON(convertToStateToSend(chessGame, i))
				}

				// check if game is ended

				// TODO: Refactor logic
				if chessGame.Winner == models.Stalemate {
					for _, client := range game.Clients {
						client.Conn.WriteMessage(1, []byte("Stalemate"))
					}
					gameOver = true
					break
				} else if chessGame.Winner == models.WhiteWins {
					for _, client := range game.Clients {
						client.Conn.WriteMessage(1, []byte("White Wins!"))
					}
					gameOver = true
					break
				} else if chessGame.Winner == models.BlackWins {
					for _, client := range game.Clients {
						client.Conn.WriteMessage(1, []byte("Black Wins!"))
					}
					gameOver = true
					break
				}
			}
		}
	}
}
