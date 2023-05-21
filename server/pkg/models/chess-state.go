package models

import (
	"fmt"
)

// TODO: castling logic

// records various information about the state of a chess position
// current board
// current turn as well as previous move played
// legality of castling for each side
type ChessState struct {
	board               *ChessBoard
	turn                int8
	previousMove        Move
	whiteCanCastleShort bool
	whiteCanCastleLong  bool
	blackCanCastleShort bool
	blackCanCastleLong  bool
}

// creates new game state
func NewChessState() *ChessState {
	return &ChessState{
		board:               NewChessBoard(),
		turn:                White,
		previousMove:        Move{},
		whiteCanCastleShort: true,
		whiteCanCastleLong:  true,
		blackCanCastleShort: true,
		blackCanCastleLong:  true,
	}
}

// returns a slice of all legal moves for a ChessState object
func (state *ChessState) EnumerateMoves() []Move {
	if state.turn == White {
		fmt.Println("Enumerating moves for white.")
		return state.enumerateMovesWhite()
	} else if state.turn == Black {
		fmt.Println("Enumerating moves for black.")
		return state.enumerateMovesBlack()
	} else {
		fmt.Println("Trying to enumerate moves for state with invalid turn value.")
		return []Move{}
	}
}

// returns a slice of all legal moves for white for a ChessState object
func (state *ChessState) enumerateMovesWhite() []Move {
	moves := make([]Move, 0, 64)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			currentPiece := state.board[i][j]
			if currentPiece <= 0 {
				continue
			} else if currentPiece == WhiteQueen {
				moves = state.enumerateMovesWhiteQueen(moves, i, j)
			} else if currentPiece == WhiteRook {
				moves = state.enumerateMovesWhiteRook(moves, i, j)
			} else if currentPiece == WhiteBishop {
				moves = state.enumerateMovesWhiteBishop(moves, i, j)
			} else if currentPiece == WhiteKnight {
				moves = state.enumerateMovesWhiteKnight(moves, i, j)
			} else if currentPiece == WhiteKing {
				moves = state.enumerateMovesWhiteKing(moves, i, j)
			} else if currentPiece == WhitePawn {
				moves = state.enumerateMovesWhitePawn(moves, i, j)
			}
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhitePawn(moves []Move, i, j int) []Move {
	if i == 1 {
		// check for single move
		if state.board[i+1][j] == EmptySquare {
			move := NewMove(Normal, i, j, i+1, j)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			// check for double moves
			if state.board[i+2][j] == EmptySquare {
				doubleMove := NewMove(Normal, i, j, i+2, j)
				if state.isLegalMove(doubleMove) {
					moves = append(moves, doubleMove)
				}
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i+1][j-1] < 0 {
			move := NewMove(Normal, i, j, i+1, j-1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := NewMove(Normal, i, j, i+1, j+1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
	} else if i == 4 {
		// check for single push
		if state.board[i+1][j] == EmptySquare {
			move := NewMove(Normal, i, j, i+1, j)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i+1][j-1] < 0 {
			move := NewMove(Normal, i, j, i+1, j-1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := NewMove(Normal, i, j, i+1, j+1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		// check for en peasant
		prev := state.previousMove
		if prev.Type == Normal {
			if prev.OldSquare.Row == 6 && prev.NewSquare.Row == 4 && state.board[prev.NewSquare.Row][prev.NewSquare.Col] == BlackPawn {
				if j-1 == prev.NewSquare.Col {
					move := NewMove(EnPassant, i, j, i+1, j-1)
					if state.isLegalMove(move) {
						moves = append(moves, move)
					}
				} else if j+1 == prev.NewSquare.Col {
					move := NewMove(EnPassant, i, j, i+1, j+1)
					if state.isLegalMove(move) {
						moves = append(moves, move)
					}
				}
			}
		}
	} else if i == 6 {
		// check for single push
		if state.board[i+1][j] == EmptySquare {
			movePromoteQueen := NewMove(PromoteQueen, i, j, i+1, j)
			movePromoteRook := NewMove(PromoteRook, i, j, i+1, j)
			movePromoteBishop := NewMove(PromoteBishop, i, j, i+1, j)
			movePromoteKnight := NewMove(PromoteKnight, i, j, i+1, j)
			if state.isLegalMove(movePromoteQueen) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i+1][j-1] < 0 {
			movePromoteQueen := NewMove(PromoteQueen, i, j, i+1, j-1)
			movePromoteRook := NewMove(PromoteRook, i, j, i+1, j-1)
			movePromoteBishop := NewMove(PromoteBishop, i, j, i+1, j-1)
			movePromoteKnight := NewMove(PromoteKnight, i, j, i+1, j-1)
			if state.isLegalMove(movePromoteQueen) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		if j+1 <= 7 && state.board[i+1][j+1] < 0 {
			movePromoteQueen := NewMove(PromoteQueen, i, j, i+1, j+1)
			movePromoteRook := NewMove(PromoteRook, i, j, i+1, j+1)
			movePromoteBishop := NewMove(PromoteBishop, i, j, i+1, j+1)
			movePromoteKnight := NewMove(PromoteKnight, i, j, i+1, j+1)
			if state.isLegalMove(movePromoteQueen) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
	} else {
		// check for single push
		if state.board[i+1][j] == EmptySquare {
			move := NewMove(Normal, i, j, i+1, j)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i+1][j-1] < 0 {
			move := NewMove(Normal, i, j, i+1, j-1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := NewMove(Normal, i, j, i+1, j+1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteKnight(moves []Move, i, j int) []Move {
	if i+1 <= 7 && j+2 <= 7 && state.board[i+1][j+2] <= 0 {
		move := NewMove(Normal, i, j, i+1, j+2)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-2 >= 0 && state.board[i+1][j-2] <= 0 {
		move := NewMove(Normal, i, j, i+1, j-2)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j+1 <= 7 && state.board[i+2][j+1] <= 0 {
		move := NewMove(Normal, i, j, i+2, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j-1 >= 0 && state.board[i+2][j-1] <= 0 {
		move := NewMove(Normal, i, j, i+2, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+2 <= 7 && state.board[i-1][j+2] <= 0 {
		move := NewMove(Normal, i, j, i-1, j+2)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-2 >= 0 && state.board[i-1][j-2] <= 0 {
		move := NewMove(Normal, i, j, i-1, j-2)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j+1 <= 7 && state.board[i-2][j+1] <= 0 {
		move := NewMove(Normal, i, j, i-2, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j-1 >= 0 && state.board[i-2][j-1] <= 0 {
		move := NewMove(Normal, i, j, i-2, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteBishop(moves []Move, i, j int) []Move {
	for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteRook(moves []Move, i, j int) []Move {
	for x, y := i, j+1; y <= 7; y++ {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j-1; y >= 0; y-- {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j; x <= 7; x++ {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j; x >= 0; x-- {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteQueen(moves []Move, i, j int) []Move {
	// bishop like movement
	for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}

	// rook like movement
	for x, y := i, j+1; y <= 7; y++ {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j-1; y >= 0; y-- {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j; x <= 7; x++ {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j; x >= 0; x-- {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteKing(moves []Move, i, j int) []Move {
	if i+1 <= 7 && j+1 <= 7 && state.board[i+1][j+1] <= 0 {
		move := NewMove(Normal, i, j, i+1, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if j+1 <= 7 && state.board[i][j+1] <= 0 {
		move := NewMove(Normal, i, j, i, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+1 <= 7 && state.board[i-1][j+1] <= 0 {
		move := NewMove(Normal, i, j, i-1, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && state.board[i+1][j] <= 0 {
		move := NewMove(Normal, i, j, i+1, j)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && state.board[i-1][j] <= 0 {
		move := NewMove(Normal, i, j, i-1, j)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-1 >= 0 && state.board[i+1][j-1] <= 0 {
		move := NewMove(Normal, i, j, i+1, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if j-1 >= 0 && state.board[i][j-1] <= 0 {
		move := NewMove(Normal, i, j, i, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-1 >= 0 && state.board[i-1][j-1] <= 0 {
		move := NewMove(Normal, i, j, i-1, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}

	// castling short
	if state.whiteCanCastleShort {
		if state.board[0][5] == EmptySquare &&
			state.board[0][6] == EmptySquare &&
			!state.board.IsSquareAttackedByBlack(0, 4) &&
			!state.board.IsSquareAttackedByBlack(0, 5) {

			move := NewMove(CastleShort, i, j, i, j+2)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
	}

	// castling long
	if state.whiteCanCastleLong {
		if state.board[0][1] == EmptySquare &&
			state.board[0][2] == EmptySquare &&
			state.board[0][3] == EmptySquare &&
			state.board.IsSquareAttackedByBlack(0, 3) &&
			state.board.IsSquareAttackedByBlack(0, 4) {

			move := NewMove(CastleLong, i, j, i, j-2)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

// returns a slice of all legal moves for black for a ChessState object
func (state *ChessState) enumerateMovesBlack() []Move {
	moves := make([]Move, 0, 64)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			currentPiece := state.board[i][j]
			if currentPiece >= 0 {
				continue
			} else if currentPiece == BlackQueen {
				moves = state.enumerateMovesBlackQueen(moves, i, j)
			} else if currentPiece == BlackRook {
				moves = state.enumerateMovesBlackRook(moves, i, j)
			} else if currentPiece == BlackBishop {
				moves = state.enumerateMovesBlackBishop(moves, i, j)
			} else if currentPiece == BlackKnight {
				moves = state.enumerateMovesBlackKnight(moves, i, j)
			} else if currentPiece == BlackKing {
				moves = state.enumerateMovesBlackKing(moves, i, j)
			} else if currentPiece == BlackPawn {
				moves = state.enumerateMovesBlackPawn(moves, i, j)
			}
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesBlackPawn(moves []Move, i, j int) []Move {
	if i == 2 {
		// check for single move
		if state.board[i-1][j] == EmptySquare {
			move := NewMove(Normal, i, j, i-1, j)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			// check for double move
			if state.board[i-2][j] == EmptySquare {
				doubleMove := NewMove(Normal, i, j, i-2, j)
				if state.isLegalMove(doubleMove) {
					moves = append(moves, doubleMove)
				}
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := NewMove(Normal, i, j, i-1, j-1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i-1][j+1] > 0 {
			move := NewMove(Normal, i, j, i-1, j+1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
	} else if i == 3 {
		// check for single push
		if state.board[i-1][j] == EmptySquare {
			move := NewMove(Normal, i, j, i-1, j)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := NewMove(Normal, i, j, i-1, j-1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i-1][j+1] > 0 {
			move := NewMove(Normal, i, j, i-1, j+1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		// check for en peasant
		prev := state.previousMove
		if prev.Type == Normal {
			if prev.OldSquare.Row == 1 && prev.NewSquare.Row == 3 && state.board[prev.NewSquare.Row][prev.NewSquare.Col] == WhitePawn {
				if j-1 == prev.NewSquare.Col {
					move := NewMove(EnPassant, i, j, i-1, j-1)
					if state.isLegalMove(move) {
						moves = append(moves, move)
					}
				} else if j+1 == prev.NewSquare.Col {
					move := NewMove(EnPassant, i, j, i-1, j+1)
					if state.isLegalMove(move) {
						moves = append(moves, move)
					}
				}
			}
		}

	} else if i == 1 {
		// check for single push
		if state.board[i-1][j] == EmptySquare {
			movePromoteQueen := NewMove(PromoteQueen, i, j, i-1, j)
			movePromoteRook := NewMove(PromoteRook, i, j, i-1, j)
			movePromoteBishop := NewMove(PromoteBishop, i, j, i-1, j)
			movePromoteKnight := NewMove(PromoteKnight, i, j, i-1, j)
			if state.isLegalMove(movePromoteQueen) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i-1][j-1] > 0 {
			movePromoteQueen := NewMove(PromoteQueen, i, j, i-1, j-1)
			movePromoteRook := NewMove(PromoteRook, i, j, i-1, j-1)
			movePromoteBishop := NewMove(PromoteBishop, i, j, i-1, j-1)
			movePromoteKnight := NewMove(PromoteKnight, i, j, i-1, j-1)
			if state.isLegalMove(movePromoteQueen) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		if j+1 <= 7 && state.board[i-1][j+1] > 0 {
			movePromoteQueen := NewMove(PromoteQueen, i, j, i-1, j+1)
			movePromoteRook := NewMove(PromoteRook, i, j, i-1, j+1)
			movePromoteBishop := NewMove(PromoteBishop, i, j, i-1, j+1)
			movePromoteKnight := NewMove(PromoteKnight, i, j, i-1, j+1)
			if state.isLegalMove(movePromoteQueen) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
	} else {
		// check for single push
		if state.board[i-1][j] == EmptySquare {
			move := NewMove(Normal, i, j, i-1, j)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := NewMove(Normal, i, j, i-1, j-1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i-1][j+1] > 0 {
			move := NewMove(Normal, i, j, i+1, j-1)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func (state *ChessState) enumerateMovesBlackKnight(moves []Move, i, j int) []Move {
	if i+1 <= 7 && j+2 <= 7 && state.board[i+1][j+2] >= 0 {
		move := NewMove(Normal, i, j, i+1, j+2)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-2 >= 0 && state.board[i+1][j-2] >= 0 {
		move := NewMove(Normal, i, j, i+1, j-2)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j+1 <= 7 && state.board[i+2][j+1] >= 0 {
		move := NewMove(Normal, i, j, i+2, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j-1 >= 0 && state.board[i+2][j-1] >= 0 {
		move := NewMove(Normal, i, j, i+2, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+2 <= 7 && state.board[i-1][j+2] >= 0 {
		move := NewMove(Normal, i, j, i-1, j+2)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-2 >= 0 && state.board[i-1][j-2] >= 0 {
		move := NewMove(Normal, i, j, i-1, j-2)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j+1 <= 7 && state.board[i-2][j+1] >= 0 {
		move := NewMove(Normal, i, j, i-2, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j-1 >= 0 && state.board[i-2][j-1] >= 0 {
		move := NewMove(Normal, i, j, i-2, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesBlackBishop(moves []Move, i, j int) []Move {
	for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesBlackRook(moves []Move, i, j int) []Move {
	for x, y := i, j+1; y <= 7; y++ {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j-1; y >= 0; y-- {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j; x <= 7; x++ {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j; x >= 0; x-- {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesBlackQueen(moves []Move, i, j int) []Move {
	for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j+1; y <= 7; y++ {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j-1; y >= 0; y-- {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j; x <= 7; x++ {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j; x >= 0; x-- {
		if state.board[x][y] == EmptySquare {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := NewMove(Normal, i, j, x, y)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesBlackKing(moves []Move, i, j int) []Move {
	// TODO: Castling
	if i+1 <= 7 && j+1 <= 7 && state.board[i+1][j+1] >= 0 {
		move := NewMove(Normal, i, j, i+1, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if j+1 <= 7 && state.board[i][j+1] >= 0 {
		move := NewMove(Normal, i, j, i, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+1 <= 7 && state.board[i-1][j+1] >= 0 {
		move := NewMove(Normal, i, j, i-1, j+1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && state.board[i+1][j] >= 0 {
		move := NewMove(Normal, i, j, i+1, j)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && state.board[i-1][j] >= 0 {
		move := NewMove(Normal, i, j, i-1, j)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-1 >= 0 && state.board[i+1][j-1] >= 0 {
		move := NewMove(Normal, i, j, i+1, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if j-1 >= 0 && state.board[i][j-1] >= 0 {
		move := NewMove(Normal, i, j, i, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-1 >= 0 && state.board[i-1][j-1] >= 0 {
		move := NewMove(Normal, i, j, i-1, j-1)
		if state.isLegalMove(move) {
			moves = append(moves, move)
		}
	}

	// castling short
	if state.blackCanCastleShort {
		if state.board[7][5] == EmptySquare &&
			state.board[7][6] == EmptySquare &&
			!state.board.IsSquareAttackedByWhite(7, 4) &&
			!state.board.IsSquareAttackedByWhite(7, 5) {

			move := NewMove(CastleShort, i, j, i, j+2)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
	}

	// castling long
	if state.blackCanCastleLong {
		if state.board[7][1] == EmptySquare &&
			state.board[7][2] == EmptySquare &&
			state.board[7][3] == EmptySquare &&
			state.board.IsSquareAttackedByWhite(7, 3) &&
			state.board.IsSquareAttackedByWhite(7, 4) {

			move := NewMove(CastleLong, i, j, i, j-2)
			if state.isLegalMove(move) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

// returns true if the move is legal
// returns false if making the move would leave self in check at the end of turn
func (state *ChessState) isLegalMove(move Move) bool {
	board, _, _, _, _ := executeMoveOnBoard(move, *state.board)
	if state.turn == White {
		return !board.IsWhiteInCheck()
	} else if state.turn == Black {
		return !board.IsBlackInCheck()
	} else {
		fmt.Println("Wait a minute... that isn't an option for whose turn it can be.")
		return false
	}
}

func (state *ChessState) ExecuteMoveOnState(move Move) *ChessState {

	newBoard, wCastleShort, wCastleLong, bCastleShort, bCastleLong := executeMoveOnBoard(move, *state.board)

	state.board = &newBoard
	state.whiteCanCastleShort = state.whiteCanCastleShort && wCastleShort
	state.whiteCanCastleLong = state.whiteCanCastleLong && wCastleLong
	state.blackCanCastleShort = state.blackCanCastleShort && bCastleShort
	state.blackCanCastleLong = state.whiteCanCastleLong && bCastleLong

	state.previousMove = move
	if state.turn == White {
		state.turn = Black
	} else if state.turn == Black {
		state.turn = White
	} else {
		fmt.Println("What is even going on???")
	}

	return state
}
