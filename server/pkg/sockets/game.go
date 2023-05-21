package sockets

import "fmt"

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

	// TODO: Setup gamestate
	currentClientCount := 0

	gameOver := false

	for !gameOver {
		select {
		case client := <-game.Register:
			game.Clients[currentClientCount] = client
			fmt.Println("Number of players in lobby: ", len(game.Clients))
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
			for _, client := range game.Clients {
				client.Conn.WriteMessage(move.Type, []byte(move.Body))
			}
		}
		// TODO: finish up connection and game logic
	}
}
