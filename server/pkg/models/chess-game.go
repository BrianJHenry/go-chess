package models

type Result string

const (
	ContinueGame Result = "C"
	Stalemate    Result = "S"
	WhiteWins    Result = "W"
	BlackWins    Result = "B"
)

type ChessGame struct {
	CurrentState  *ChessState
	MoveHistory   []Move
	PossibleMoves []Move
	Winner        Result
}

func NewChessGame() ChessGame {
	game := ChessGame{
		CurrentState: NewChessState(),
		MoveHistory:  make([]Move, 0, 64),
		Winner:       ContinueGame,
	}
	game.PossibleMoves = game.CurrentState.EnumerateMoves()
	return game
}

func (game *ChessGame) ExecuteMoveOnGame(move Move) {
	game.MoveHistory = append(game.MoveHistory, move)
	game.CurrentState = game.CurrentState.ExecuteMoveOnState(move)
	game.PossibleMoves = game.CurrentState.EnumerateMoves()
	if len(game.PossibleMoves) == 0 {
		if game.CurrentState.Turn == White {
			if game.CurrentState.Board.IsWhiteInCheck() {
				game.Winner = BlackWins
			} else {
				game.Winner = Stalemate
			}
		} else if game.CurrentState.Turn == Black {
			if game.CurrentState.Board.IsBlackInCheck() {
				game.Winner = WhiteWins
			} else {
				game.Winner = Stalemate
			}
		}
	}
}
