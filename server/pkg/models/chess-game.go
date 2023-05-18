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

func (game *ChessGame) ExecuteMove(move Move) {
	game.Moves = append(game.Moves, move)
	game.CurrentState = game.CurrentState.progressStateForward(move)
	game.CurrentState.board.printChessBoard()
}