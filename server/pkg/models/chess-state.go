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
		previousMove:        0,
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
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			// check for double moves
			if state.board[i+2][j] == EmptySquare {
				doubleMove := encodeMove(state.turn, NoSpecial, i, j, i+2, j)
				if !moveAbandonsKing(state.board, doubleMove, state.turn) {
					moves = append(moves, doubleMove)
				}
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i+1][j-1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
	} else if i == 4 {
		// check for single push
		if state.board[i+1][j] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i+1][j-1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for en peasant
		prev := uint16(state.previousMove)
		special := (prev & SpecialMask) >> 13
		if special == 0 {
			i1 := (prev & StartIMask) >> 1
			i2 := (prev & EndIMask) >> 7
			j2 := (prev & EndJMask) >> 10
			if i1 == 6 && i2 == 4 && state.board[i2][j2] == BlackPawn {
				if j-1 == int(j2) {
					move := encodeMove(state.turn, EnPeasant, i, j, i+1, j-1)
					if !moveAbandonsKing(state.board, move, state.turn) {
						moves = append(moves, move)
					}
				} else if j+1 == int(j2) {
					move := encodeMove(state.turn, EnPeasant, i, j, i+1, j+1)
					if !moveAbandonsKing(state.board, move, state.turn) {
						moves = append(moves, move)
					}
				}
			}
		}
	} else if i == 6 {
		// check for single push
		if state.board[i+1][j] == EmptySquare {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i+1, j)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i+1, j)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i+1, j)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i+1, j)
			if !moveAbandonsKing(state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i+1][j-1] < 0 {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i+1, j-1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i+1, j-1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i+1, j-1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i+1, j-1)
			if !moveAbandonsKing(state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		if j+1 <= 7 && state.board[i+1][j+1] < 0 {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i+1, j+1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i+1, j+1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i+1, j+1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i+1, j+1)
			if !moveAbandonsKing(state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
	} else {
		// check for single push
		if state.board[i+1][j] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i+1][j-1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteKnight(moves []Move, i, j int) []Move {
	if i+1 <= 7 && j+2 <= 7 && state.board[i+1][j+2] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+2)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-2 >= 0 && state.board[i+1][j-2] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-2)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j+1 <= 7 && state.board[i+2][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+2, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j-1 >= 0 && state.board[i+2][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+2, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+2 <= 7 && state.board[i-1][j+2] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+2)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-2 >= 0 && state.board[i-1][j-2] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-2)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j+1 <= 7 && state.board[i-2][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-2, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j-1 >= 0 && state.board[i-2][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-2, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteBishop(moves []Move, i, j int) []Move {
	for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
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
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j-1; y >= 0; y-- {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j; x <= 7; x++ {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j; x >= 0; x-- {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
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
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
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
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j-1; y >= 0; y-- {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j; x <= 7; x++ {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j; x >= 0; x-- {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
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
	// TODO: Castling
	if i+1 <= 7 && j+1 <= 7 && state.board[i+1][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if j+1 <= 7 && state.board[i][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+1 <= 7 && state.board[i-1][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && state.board[i+1][j] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && state.board[i-1][j] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-1 >= 0 && state.board[i+1][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if j-1 >= 0 && state.board[i][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-1 >= 0 && state.board[i-1][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
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
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			// check for double move
			if state.board[i-2][j] == EmptySquare {
				doubleMove := encodeMove(state.turn, NoSpecial, i, j, i-2, j)
				if !moveAbandonsKing(state.board, doubleMove, state.turn) {
					moves = append(moves, doubleMove)
				}
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i-1][j+1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
	} else if i == 3 {
		// check for single push
		if state.board[i-1][j] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i-1][j+1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for en peasant
		prev := uint16(state.previousMove)
		special := (prev & SpecialMask) >> 13
		if special == 0 {
			i1 := (prev & StartJMask) >> 1
			i2 := (prev & EndIMask) >> 7
			j2 := (prev & EndJMask) >> 10
			if i1 == 1 && i2 == 3 && state.board[i2][j2] == WhitePawn {
				if j-1 == int(j2) {
					move := encodeMove(state.turn, EnPeasant, i, j, i-1, j-1)
					if !moveAbandonsKing(state.board, move, state.turn) {
						moves = append(moves, move)
					}
				} else if j+1 == int(j2) {
					move := encodeMove(state.turn, EnPeasant, i, j, i-1, j+1)
					if !moveAbandonsKing(state.board, move, state.turn) {
						moves = append(moves, move)
					}
				}
			}
		}

	} else if i == 1 {
		// check for single push
		if state.board[i-1][j] == EmptySquare {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i-1, j)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i-1, j)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i-1, j)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i-1, j)
			if !moveAbandonsKing(state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i-1][j-1] > 0 {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i-1, j-1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i-1, j-1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i-1, j-1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i-1, j-1)
			if !moveAbandonsKing(state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		if j+1 <= 7 && state.board[i-1][j+1] > 0 {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i-1, j+1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i-1, j+1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i-1, j+1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i-1, j+1)
			if !moveAbandonsKing(state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
	} else {
		// check for single push
		if state.board[i-1][j] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if j-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if j+1 <= 7 && state.board[i-1][j+1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func (state *ChessState) enumerateMovesBlackKnight(moves []Move, i, j int) []Move {
	if i+1 <= 7 && j+2 <= 7 && state.board[i+1][j+2] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+2)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-2 >= 0 && state.board[i+1][j-2] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-2)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j+1 <= 7 && state.board[i+2][j+1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+2, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j-1 >= 0 && state.board[i+2][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+2, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+2 <= 7 && state.board[i-1][j+2] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+2)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-2 >= 0 && state.board[i-1][j-2] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-2)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j+1 <= 7 && state.board[i-2][j+1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-2, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j-1 >= 0 && state.board[i-2][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-2, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesBlackBishop(moves []Move, i, j int) []Move {
	for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
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
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j-1; y >= 0; y-- {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j; x <= 7; x++ {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j; x >= 0; x-- {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
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
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j+1; y <= 7; y++ {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i, j-1; y >= 0; y-- {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i+1, j; x <= 7; x++ {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
			break
		} else {
			break
		}
	}
	for x, y := i-1, j; x >= 0; x-- {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(state.board, move, state.turn) {
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
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if j+1 <= 7 && state.board[i][j+1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+1 <= 7 && state.board[i-1][j+1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && state.board[i+1][j] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && state.board[i-1][j] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-1 >= 0 && state.board[i+1][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if j-1 >= 0 && state.board[i][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-1 >= 0 && state.board[i-1][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
		if !moveAbandonsKing(state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	return moves
}

// returns true if move illegally places king in check
func moveAbandonsKing(board *ChessBoard, move Move, turn int8) bool {
	newBoard := testMove(move, *board)
	if turn == White {
		return newBoard.WhiteInCheck()
	} else {
		return newBoard.BlackInCheck()
	}
}

func (state *ChessState) progressStateForward(move Move) *ChessState {
	state.previousMove = move
	turn, special, i1, j1, i2, j2 := decodeMove(move)

	// check castling	4-7 i
	if special == CastleShort {
		if turn == White {
			state.board[4][0] = EmptySquare
			state.board[5][0] = WhiteRook
			state.board[6][0] = WhiteKing
			state.board[7][0] = EmptySquare
			state.whiteCanCastleShort = false
			state.whiteCanCastleLong = false
		} else {
			state.board[4][7] = EmptySquare
			state.board[5][7] = WhiteRook
			state.board[6][7] = WhiteKing
			state.board[7][7] = EmptySquare
			state.blackCanCastleShort = false
			state.blackCanCastleLong = false
		}
		state.turn = int8(turn+1) % 2
		return state
	} else if special == CastleLong {
		if turn == White {
			state.board[0][0] = EmptySquare
			state.board[1][0] = EmptySquare
			state.board[2][0] = WhiteKing
			state.board[3][0] = WhiteRook
			state.board[4][0] = EmptySquare
			state.whiteCanCastleShort = false
			state.whiteCanCastleLong = false
		} else {
			state.board[0][7] = EmptySquare
			state.board[1][7] = EmptySquare
			state.board[2][7] = BlackKing
			state.board[3][7] = BlackRook
			state.board[4][7] = EmptySquare
			state.blackCanCastleShort = false
			state.blackCanCastleLong = false
		}
		state.turn = int8(turn+1) % 2
		return state
	}
	var movingPieceType int8 = 0
	if special == NoSpecial || special == EnPeasant {
		movingPieceType = state.board[i1][j1]
		if movingPieceType == WhiteKing {
			state.whiteCanCastleLong = false
			state.whiteCanCastleShort = false
		} else if movingPieceType == BlackKing {
			state.blackCanCastleLong = false
			state.blackCanCastleShort = false
		}
	} else if special == PromoteToQueen {
		if turn == White {
			movingPieceType = WhiteQueen
		} else {
			movingPieceType = BlackQueen
		}
	} else if special == PromoteToRook {
		if turn == White {
			movingPieceType = WhiteRook
		} else {
			movingPieceType = BlackRook
		}
	} else if special == PromoteToBishop {
		if turn == White {
			movingPieceType = WhiteBishop
		} else {
			movingPieceType = BlackBishop
		}
	} else if special == PromoteToKnight {
		if turn == White {
			movingPieceType = WhiteKnight
		} else {
			movingPieceType = BlackKnight
		}
	}

	if movingPieceType == WhiteRook && i1 == 0 && j1 == 0 {
		state.whiteCanCastleLong = false
	} else if movingPieceType == WhiteRook && i1 == 0 && j1 == 7 {
		state.whiteCanCastleShort = false
	} else if movingPieceType == BlackRook && i1 == 7 && j1 == 0 {
		state.blackCanCastleLong = false
	} else if movingPieceType == BlackRook && i1 == 7 && j1 == 7 {
		state.blackCanCastleShort = false
	}

	state.board[i1][j1] = EmptySquare
	state.board[i2][j2] = movingPieceType
	if special == EnPeasant {
		state.board[i2][j1] = EmptySquare
	}
	state.turn = int8(turn+1) % 2
	return state
}

func (state *ChessState) progressStateBackward(move Move) {

}
