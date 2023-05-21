package chesssockets

import "fmt"

type Game struct {
	// game info
	NumberOfPlayers int

	Clients []*Client

	// websocket handling
	Register    chan *Client
	RecieveMove chan Message
}

func NewGame(numberOfPlayers int) *Game {
	return &Game{
		NumberOfPlayers: numberOfPlayers,
		Clients:         make([]*Client, numberOfPlayers),
		Register:        make(chan *Client),
		RecieveMove:     make(chan Message),
	}
}

func (game *Game) Start() {

	// TODO: Setup gamestate
	currentClientCount := 0

	for {
		select {
		case client := <-game.Register:
			game.Clients[currentClientCount] = client
			fmt.Println("Number of players in lobby: ", len(game.Clients))
		case move := <-game.RecieveMove:
			for _, client := range game.Clients {
				client.Conn.WriteMessage(move.Type, []byte(move.Body))
			}
		}
		// TODO: finish up connection and game logic
	}
}
