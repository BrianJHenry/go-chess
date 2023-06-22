package sockets

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/BrianJHenry/go-chess/server/pkg/models"
)

// setup mappings for move types
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

// setup moves to be sent across websockets
type APIMove struct {
	MoveType  string `json:"moveType"`
	OldSquare int    `json:"oldSquare"`
	NewSquare int    `json:"newSquare"`
}

func convertToAPIMove(move models.Move) APIMove {
	moveType := moveTypesArray[move.Type]
	oldSquare := move.OldSquare.Row*8 + move.OldSquare.Col
	newSquare := move.NewSquare.Row*8 + move.NewSquare.Col

	return APIMove{
		MoveType:  moveType,
		OldSquare: oldSquare,
		NewSquare: newSquare,
	}
}

func convertToMove(move APIMove) models.Move {
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
type APIState struct {
	Turn          bool      `json:"turn"`
	Board         []int8    `json:"board"`
	PreviousMoves []APIMove `json:"previousMoves"`
	PossibleMoves []APIMove `json:"possibleMoves"`
}

func convertToAPIState(game models.ChessGame, ownColor int) APIState {
	var board = make([]int8, 0, 64)
	for _, row := range game.CurrentState.Board {
		sliceRow := row[:]
		board = append(board, sliceRow...)
	}
	var turn bool
	var possibleMoves []APIMove
	log.Printf("Own: %v : Current Turn: %v", ownColor, game.CurrentState.Turn)
	if ownColor == int(game.CurrentState.Turn) {
		turn = true
		for _, move := range game.PossibleMoves {
			possibleMoves = append(possibleMoves, convertToAPIMove(move))
		}
	} else {
		turn = false
		possibleMoves = make([]APIMove, 0)
	}
	var previousMoves = make([]APIMove, 0, len(game.MoveHistory))
	for _, move := range game.MoveHistory {
		convertedMove := convertToAPIMove(move)
		previousMoves = append(previousMoves, convertedMove)
	}

	return APIState{
		Turn:          turn,
		Board:         board,
		PreviousMoves: previousMoves,
		PossibleMoves: possibleMoves,
	}
}

// empty game state
func CreateEmptyGameState() APIState {
	return APIState{
		Turn:  false,
		Board: []int8{-99},
		PreviousMoves: []APIMove{
			{
				MoveType:  "",
				NewSquare: -99,
				OldSquare: -99,
			},
		},
		PossibleMoves: []APIMove{
			{
				MoveType:  "",
				NewSquare: -99,
				OldSquare: -99,
			},
		},
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
	RecieveMove chan APIMove
}

func NewGame(numberOfPlayers int, gameID string, delete func(id string)) *Game {
	return &Game{
		GameID:          gameID,
		Delete:          delete,
		NumberOfPlayers: numberOfPlayers,
		Clients:         make([]*Client, 0, numberOfPlayers),
		Register:        make(chan *Client),
		Unregister:      make(chan *Client),
		RecieveMove:     make(chan APIMove),
	}
}

func (game *Game) Start() {
	defer func() {
		game.Delete(game.GameID)
	}()

	// create new game
	chessGame := models.NewChessGame()
	gameOver := false

	var message Message

	for !gameOver {
		select {
		case client := <-game.Register:
			log.Println("Doing register work...")
			game.Clients = append(game.Clients, client)
			// if we have enough players start the game
			if len(game.Clients) == game.NumberOfPlayers {
				for i, c := range game.Clients {
					// send message
					message = NewMessage(UpdateStateMessage, "", convertToAPIState(chessGame, i))
					log.Println(message)
					// TODO: Error checking
					c.Conn.WriteJSON(message)
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

					// send message
					message = NewMessage(MiscMessage, "Opponent Disconnected", CreateEmptyGameState())
					// TODO: Error checking
					c.Conn.WriteJSON(message)

					message = NewMessage(GameInfoMessage, fmt.Sprintf("%s wins!", winner), CreateEmptyGameState())
					// TODO: Error checking
					c.Conn.WriteJSON(message)
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
				message = NewMessage(MiscMessage, "Invalid Move.", CreateEmptyGameState())
				// TODO: Error checking
				for _, client := range game.Clients {
					client.Conn.WriteJSON(message)
				}
				break
			}

			// execute move
			chessGame.ExecuteMoveOnGame(tryMove)

			// send back updated state
			for i, client := range game.Clients {
				log.Println("Move recieved.")
				// send updated state
				message = NewMessage(UpdateStateMessage, "", convertToAPIState(chessGame, i))
				client.Conn.WriteJSON(message)
			}

			// check if game is ended

			// TODO: Refactor logic
			if chessGame.Winner == models.Stalemate {
				for _, client := range game.Clients {
					// send message
					message = NewMessage(GameInfoMessage, "Stalemate", CreateEmptyGameState())
					client.Conn.WriteJSON(message)
				}
				gameOver = true
				break
			} else if chessGame.Winner == models.WhiteWins {
				for _, client := range game.Clients {
					// send message
					message = NewMessage(GameInfoMessage, "White wins!", CreateEmptyGameState())
					client.Conn.WriteJSON(message)
				}
				gameOver = true
				break
			} else if chessGame.Winner == models.BlackWins {
				for _, client := range game.Clients {
					// send message
					message = NewMessage(GameInfoMessage, "Black wins!", CreateEmptyGameState())
					client.Conn.WriteJSON(message)
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
					// send message
					message = NewMessage(UpdateStateMessage, "", convertToAPIState(chessGame, i))
					client.Conn.WriteJSON(message)
				}

				// check if game is ended

				// TODO: Refactor logic
				if chessGame.Winner == models.Stalemate {
					for _, client := range game.Clients {
						// send message
						message = NewMessage(GameInfoMessage, "Stalemate", CreateEmptyGameState())
						client.Conn.WriteJSON(message)
					}
					gameOver = true
					break
				} else if chessGame.Winner == models.WhiteWins {
					for _, client := range game.Clients {
						// send message
						message = NewMessage(GameInfoMessage, "White wins!", CreateEmptyGameState())
						client.Conn.WriteJSON(message)
					}
					gameOver = true
					break
				} else if chessGame.Winner == models.BlackWins {
					for _, client := range game.Clients {
						// send message
						message = NewMessage(GameInfoMessage, "Black wins!", CreateEmptyGameState())
						client.Conn.WriteJSON(message)
					}
					gameOver = true
					break
				}
			}
		}
	}
}
