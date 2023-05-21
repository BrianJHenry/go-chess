package models

type ChessGame struct {
	CurrentState *ChessState
	Moves        []Move
}

func NewChessGame() *ChessGame {
	return &ChessGame{
		CurrentState: NewChessState(),
		Moves:        make([]Move, 0, 64),
	}
}

func (game *ChessGame) ExecuteMoveOnGame(move Move) {
	game.Moves = append(game.Moves, move)
	game.CurrentState = game.CurrentState.ExecuteMoveOnState(move)
}
