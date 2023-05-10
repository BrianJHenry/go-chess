package models

import (
	"fmt"
)

// --------------------------------------------------
// constants
// --------------------------------------------------

const (
	White = iota
	Black
)

const (
	BlackKing = iota - 6
	BlackQueen
	BlackRook
	BlackBishop
	BlackKnight
	BlackPawn
	EmptySquare
	WhitePawn
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
)

const (
	NoSpecial = iota
	CastleShort
	CastleLong
	PromoteToQueen
	PromoteToRook
	PromoteToBishop
	PromoteToKnight
	EnPeasant
)

// masks
const (
	TurnMask    = 0b0000000000000001
	StartIMask  = 0b0000000000001110
	StartJMask  = 0b0000000001110000
	EndIMask    = 0b0000001110000000
	EndJMask    = 0b0001110000000000
	SpecialMask = 0b1110000000000000
)

// --------------------------------------------------
// types
// --------------------------------------------------

// represents a bit encoding of a chess move
// bit 0 represents color making move: 0 for white; 1 for black
// bits 1-3 represent a starting i index on a board
// bits 4-6 represent a starting j index on a board
// bits 7-9 represent an ending i index on a board
// bits 10-12 represent an ending j index on a board
// bits 13-15 represent special movements
// 0 = no special move
// 1 = castle short
// 2 = castle long
// 3 = pawn promote to queen
// 4 = pawn promote to rook
// 5 = pawn promote to bishop
// 6 = pawn promote to knight
type Move uint16

// ----------------------------+
//
//	|
//
// 7	BR BN BB BQ BK BB BN BR	|
// 6	BP BP BP BP BP BP BP BP	|
// 5	-- -- -- -- -- -- -- --	|
// 4	-- -- -- -- -- -- -- --	|
// 3	-- -- -- -- -- -- -- --	|
// 2	-- -- -- -- -- -- -- --	|
// 1	WP WP WP WP WP WP WP WP	|
// 0	WR WN WB WQ WK WB WN WR	|
// j							|
//
//	i   0  1  2  3  4  5  6  7	|
//
// ----------------------------+
type ChessBoard [8][8]int8

// records various information about the state of a chess position
// current board
// current turn as well as previous move played
// legality of castling for each side
type ChessState struct {
	board               ChessBoard
	turn                int8
	previousMove        Move
	whiteCanCastleShort bool
	whiteCanCastleLong  bool
	blackCanCastleShort bool
	blackCanCastleLong  bool
}

// --------------------------------------------------
// functions/methods
// --------------------------------------------------

func StartingState() ChessState {
	var startingState ChessState = ChessState{}
	startingState.board = ChessBoard{
		{WhiteRook, WhitePawn, EmptySquare, EmptySquare, EmptySquare, EmptySquare, BlackPawn, BlackRook},
		{WhiteKnight, WhitePawn, EmptySquare, EmptySquare, EmptySquare, EmptySquare, BlackPawn, BlackKnight},
		{WhiteBishop, WhitePawn, EmptySquare, EmptySquare, EmptySquare, EmptySquare, BlackPawn, BlackBishop},
		{WhiteQueen, WhitePawn, EmptySquare, EmptySquare, EmptySquare, EmptySquare, BlackPawn, BlackQueen},
		{WhiteKing, WhitePawn, EmptySquare, EmptySquare, EmptySquare, EmptySquare, BlackPawn, BlackKing},
		{WhiteBishop, WhitePawn, EmptySquare, EmptySquare, EmptySquare, EmptySquare, BlackPawn, BlackBishop},
		{WhiteKnight, WhitePawn, EmptySquare, EmptySquare, EmptySquare, EmptySquare, BlackPawn, BlackKnight},
		{WhiteRook, WhitePawn, EmptySquare, EmptySquare, EmptySquare, EmptySquare, BlackPawn, BlackRook},
	}
	startingState.turn = White
	startingState.previousMove = 0
	startingState.whiteCanCastleShort = true
	startingState.whiteCanCastleLong = true
	startingState.blackCanCastleShort = true
	startingState.blackCanCastleLong = true
	return startingState
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
	if j == 1 {
		// check for single move
		if state.board[i][j+1] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
			// check for double moves
			if state.board[i][j+2] == EmptySquare {
				doubleMove := encodeMove(state.turn, NoSpecial, i, j, i, j+2)
				if !moveAbandonsKing(&state.board, doubleMove, state.turn) {
					moves = append(moves, doubleMove)
				}
			}
		}
		// check for captures
		if i-1 >= 0 && state.board[i-1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if i+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
	} else if j == 4 {
		// check for single push
		if state.board[i][j+1] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if i-1 >= 0 && state.board[i-1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if i+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for en peasant
		prev := uint16(state.previousMove)
		special := (prev & SpecialMask) >> 13
		if special == 0 {
			j1 := (prev & StartJMask) >> 4
			i2 := (prev & EndIMask) >> 7
			j2 := (prev & EndJMask) >> 10
			if j1 == 6 && j2 == 4 && state.board[i2][j2] == BlackPawn {
				if i-1 == int(i2) {
					move := encodeMove(state.turn, EnPeasant, i, j, i-1, j+1)
					if !moveAbandonsKing(&state.board, move, state.turn) {
						moves = append(moves, move)
					}
				} else if i+1 == int(i2) {
					move := encodeMove(state.turn, EnPeasant, i, j, i+1, j+1)
					if !moveAbandonsKing(&state.board, move, state.turn) {
						moves = append(moves, move)
					}
				}
			}
		}
	} else if j == 6 {
		// check for single push
		if state.board[i][j+1] == EmptySquare {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i, j+1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i, j+1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i, j+1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i, j+1)
			if !moveAbandonsKing(&state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		// check for captures
		if i-1 >= 0 && state.board[i-1][j+1] < 0 {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i-1, j+1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i-1, j+1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i-1, j+1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i-1, j+1)
			if !moveAbandonsKing(&state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		if i+1 <= 7 && state.board[i+1][j+1] < 0 {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i+1, j+1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i+1, j+1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i+1, j+1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i+1, j+1)
			if !moveAbandonsKing(&state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
	} else {
		// check for single push
		if state.board[i][j+1] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if i-1 >= 0 && state.board[i-1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if i+1 <= 7 && state.board[i+1][j+1] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteKnight(moves []Move, i, j int) []Move {
	if i+1 <= 7 && j+2 <= 7 && state.board[i+1][j+2] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+2)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-2 >= 0 && state.board[i+1][j-2] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-2)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j+1 <= 7 && state.board[i+2][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+2, j+1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j-1 >= 0 && state.board[i+2][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+2, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+2 <= 7 && state.board[i-1][j+2] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+2)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-2 >= 0 && state.board[i-1][j-2] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-2)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j+1 <= 7 && state.board[i-2][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-2, j+1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j-1 >= 0 && state.board[i-2][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-2, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesWhiteBishop(moves []Move, i, j int) []Move {
	for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] < 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if j+1 <= 7 && state.board[i][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i, j+1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+1 <= 7 && state.board[i-1][j+1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && state.board[i+1][j] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && state.board[i-1][j] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 >= 0 && j-1 >= 0 && state.board[i+1][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if j-1 >= 0 && state.board[i][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-1 >= 0 && state.board[i-1][j-1] <= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
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
	if j == 2 {
		// check for single move
		if state.board[i][j-1] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
			// check for double move
			if state.board[i][j-2] == EmptySquare {
				doubleMove := encodeMove(state.turn, NoSpecial, i, j, i, j-2)
				if !moveAbandonsKing(&state.board, doubleMove, state.turn) {
					moves = append(moves, doubleMove)
				}
			}
		}
		// check for captures
		if i-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if i+1 <= 7 && state.board[i+1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
	} else if j == 3 {
		// check for single push
		if state.board[i][j-1] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if i-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if i+1 <= 7 && state.board[i+1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for en peasant
		prev := uint16(state.previousMove)
		special := (prev & SpecialMask) >> 13
		if special == 0 {
			j1 := (prev & StartJMask) >> 4
			i2 := (prev & EndIMask) >> 7
			j2 := (prev & EndJMask) >> 10
			if j1 == 1 && j2 == 3 && state.board[i2][j2] == WhitePawn {
				if i-1 == int(i2) {
					move := encodeMove(state.turn, EnPeasant, i, j, i-1, j-1)
					if !moveAbandonsKing(&state.board, move, state.turn) {
						moves = append(moves, move)
					}
				} else if i+1 == int(i2) {
					move := encodeMove(state.turn, EnPeasant, i, j, i+1, j-1)
					if !moveAbandonsKing(&state.board, move, state.turn) {
						moves = append(moves, move)
					}
				}
			}
		}

	} else if j == 1 {
		// check for single push
		if state.board[i][j-1] == EmptySquare {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i, j-1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i, j-1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i, j-1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i, j-1)
			if !moveAbandonsKing(&state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		// check for captures
		if i-1 >= 0 && state.board[i-1][j-1] > 0 {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i-1, j-1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i-1, j-1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i-1, j-1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i-1, j-1)
			if !moveAbandonsKing(&state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
		if i+1 <= 7 && state.board[i+1][j-1] > 0 {
			movePromoteQueen := encodeMove(state.turn, PromoteToQueen, i, j, i+1, j-1)
			movePromoteRook := encodeMove(state.turn, PromoteToRook, i, j, i+1, j-1)
			movePromoteBishop := encodeMove(state.turn, PromoteToBishop, i, j, i+1, j-1)
			movePromoteKnight := encodeMove(state.turn, PromoteToKnight, i, j, i+1, j-1)
			if !moveAbandonsKing(&state.board, movePromoteQueen, state.turn) {
				moves = append(moves, movePromoteQueen, movePromoteRook, movePromoteBishop, movePromoteKnight)
			}
		}
	} else {
		// check for single push
		if state.board[i][j-1] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, i, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		// check for captures
		if i-1 >= 0 && state.board[i-1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
		if i+1 <= 7 && state.board[i+1][j-1] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		}
	}

	return moves
}

func (state *ChessState) enumerateMovesBlackKnight(moves []Move, i, j int) []Move {
	if i+1 <= 7 && j+2 <= 7 && state.board[i+1][j+2] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j+2)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && j-2 >= 0 && state.board[i+1][j-2] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-2)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j+1 <= 7 && state.board[i+2][j+1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+2, j+1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+2 <= 7 && j-1 >= 0 && state.board[i+2][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+2, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+2 <= 7 && state.board[i-1][j+2] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+2)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-2 >= 0 && state.board[i-1][j-2] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-2)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j+1 <= 7 && state.board[i-2][j+1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-2, j+1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-2 >= 0 && j-1 >= 0 && state.board[i-2][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-2, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (state *ChessState) enumerateMovesBlackBishop(moves []Move, i, j int) []Move {
	for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
		if state.board[x][y] == EmptySquare {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
			if !moveAbandonsKing(&state.board, move, state.turn) {
				moves = append(moves, move)
			}
		} else if state.board[x][y] > 0 {
			move := encodeMove(state.turn, NoSpecial, i, j, x, y)
			if !moveAbandonsKing(&state.board, move, state.turn) {
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
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if j+1 <= 7 && state.board[i][j+1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i, j+1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j+1 <= 7 && state.board[i-1][j+1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j+1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 <= 7 && state.board[i+1][j] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && state.board[i-1][j] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i+1 >= 0 && j-1 >= 0 && state.board[i+1][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i+1, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if j-1 >= 0 && state.board[i][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	if i-1 >= 0 && j-1 >= 0 && state.board[i-1][j-1] >= 0 {
		move := encodeMove(state.turn, NoSpecial, i, j, i-1, j-1)
		if !moveAbandonsKing(&state.board, move, state.turn) {
			moves = append(moves, move)
		}
	}
	return moves
}

func (board *ChessBoard) WhiteInCheck() bool {
	// find king
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			if board[i][j] == WhiteKing {
				// check for pawn checks
				if i-1 >= 0 && board[i-1][j+1] == BlackPawn {
					return true
				}
				if i+1 <= 7 && board[i+1][j+1] == BlackPawn {
					return true
				}
				// check for knight checks
				if i+1 <= 7 && j+2 <= 7 && board[i+1][j+2] == BlackKnight {
					return true
				}
				if i+1 <= 7 && j-2 >= 0 && board[i+1][j-2] == BlackKnight {
					return true
				}
				if i+2 <= 7 && j+1 <= 7 && board[i+2][j+1] == BlackKnight {
					return true
				}
				if i+2 <= 7 && j-1 >= 0 && board[i+2][j-1] == BlackKnight {
					return true
				}
				if i-1 >= 0 && j+2 <= 7 && board[i-1][j+2] == BlackKnight {
					return true
				}
				if i-1 >= 0 && j-2 >= 0 && board[i-1][j-2] == BlackKnight {
					return true
				}
				if i-2 >= 0 && j+1 <= 7 && board[i-2][j+1] == BlackKnight {
					return true
				}
				if i-2 >= 0 && j-1 >= 0 && board[i-2][j-1] == BlackKnight {
					return true
				}
				// check for bishop/queen checks
				for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == BlackBishop || board[x][y] == BlackQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == BlackBishop || board[x][y] == BlackQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == BlackBishop || board[x][y] == BlackQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == BlackBishop || board[x][y] == BlackQueen {
						return true
					} else {
						break
					}
				}

				// check for rook/queen checks
				for x, y := i, j+1; y <= 7; y++ {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == BlackRook || board[x][y] == BlackQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i, j-1; y >= 0; y-- {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == BlackRook || board[x][y] == BlackQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i+1, j; x <= 7; x++ {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == BlackRook || board[x][y] == BlackQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i-1, j; x >= 0; x-- {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == BlackRook || board[x][y] == BlackQueen {
						return true
					} else {
						break
					}
				}
				return false
			}
		}
	}
	// huh
	return false
}

func (board *ChessBoard) BlackInCheck() bool {
	// find king
	for j := 7; j >= 0; j-- {
		for i := 7; i >= 0; i-- {
			if board[i][j] == BlackKing {
				// check for pawn checks
				if i-1 >= 0 && board[i-1][j-1] == WhitePawn {
					return true
				}
				if i+1 <= 7 && board[i+1][j-1] == WhitePawn {
					return true
				}
				// check for knight checks
				if i+1 <= 7 && j+2 <= 7 && board[i+1][j+2] == WhiteKnight {
					return true
				}
				if i+1 <= 7 && j-2 >= 0 && board[i+1][j-2] == WhiteKnight {
					return true
				}
				if i+2 <= 7 && j+1 <= 7 && board[i+2][j+1] == WhiteKnight {
					return true
				}
				if i+2 <= 7 && j-1 >= 0 && board[i+2][j-1] == WhiteKnight {
					return true
				}
				if i-1 >= 0 && j+2 <= 7 && board[i-1][j+2] == WhiteKnight {
					return true
				}
				if i-1 >= 0 && j-2 >= 0 && board[i-1][j-2] == WhiteKnight {
					return true
				}
				if i-2 >= 0 && j+1 <= 7 && board[i-2][j+1] == WhiteKnight {
					return true
				}
				if i-2 >= 0 && j-1 >= 0 && board[i-2][j-1] == WhiteKnight {
					return true
				}
				// check for bishop/queen checks
				for x, y := i+1, j+1; x <= 7 && y <= 7; x, y = x+1, y+1 {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == WhiteBishop || board[x][y] == WhiteQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i+1, j-1; x <= 7 && y >= 0; x, y = x+1, y-1 {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == WhiteBishop || board[x][y] == WhiteQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i-1, j+1; x >= 0 && y <= 7; x, y = x-1, y+1 {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == WhiteBishop || board[x][y] == WhiteQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i-1, j-1; x >= 0 && y >= 0; x, y = x-1, y-1 {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == WhiteBishop || board[x][y] == WhiteQueen {
						return true
					} else {
						break
					}
				}

				// check for rook/queen checks
				for x, y := i, j+1; y <= 7; y++ {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == WhiteRook || board[x][y] == WhiteQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i, j-1; y >= 0; y-- {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == WhiteRook || board[x][y] == WhiteQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i+1, j; x <= 7; x++ {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == WhiteRook || board[x][y] == WhiteQueen {
						return true
					} else {
						break
					}
				}
				for x, y := i-1, j; x >= 0; x-- {
					if board[x][y] == EmptySquare {
						continue
					} else if board[x][y] == WhiteRook || board[x][y] == WhiteQueen {
						return true
					} else {
						break
					}
				}
				return false
			}
		}
	}
	// huh
	return false
}

func encodeMove(turn, specialMove int8, i1, j1, i2, j2 int) Move {
	var move uint16 = 0
	move = move | uint16(turn)
	move = move | (uint16(i1) << 1)
	move = move | (uint16(j1) << 4)
	move = move | (uint16(i2) << 7)
	move = move | (uint16(j2) << 10)
	move = move | (uint16(specialMove) << 13)
	return Move(move)
}

func doMove(move Move, board ChessBoard) ChessBoard {
	moveInfo := uint16(move)
	turn := (TurnMask & moveInfo)
	special := (SpecialMask & moveInfo) >> 13

	// check castling	4-7 i
	if special == CastleShort {
		if turn == White {
			board[4][0] = EmptySquare
			board[5][0] = WhiteRook
			board[6][0] = WhiteKing
			board[7][0] = EmptySquare
		} else {
			board[4][7] = EmptySquare
			board[5][7] = WhiteRook
			board[6][7] = WhiteKing
			board[7][7] = EmptySquare
		}
		return board
	} else if special == CastleLong {
		if turn == White {
			board[0][0] = EmptySquare
			board[1][0] = EmptySquare
			board[2][0] = WhiteKing
			board[3][0] = WhiteRook
			board[4][0] = EmptySquare
		} else {
			board[0][7] = EmptySquare
			board[1][7] = EmptySquare
			board[2][7] = BlackKing
			board[3][7] = BlackRook
			board[4][7] = EmptySquare
		}
		return board
	}

	// get move by indices
	i1 := (StartIMask & moveInfo) >> 1
	j1 := (StartJMask & moveInfo) >> 4
	i2 := (EndIMask & moveInfo) >> 7
	j2 := (EndJMask & moveInfo) >> 10

	var movingPieceType int8 = 0
	if special == NoSpecial || special == EnPeasant {
		movingPieceType = board[i1][j1]
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

	board[i1][j1] = EmptySquare
	board[i2][j2] = movingPieceType
	if special == EnPeasant {
		board[i2][j1] = EmptySquare
	}
	return board
}

// returns true if move illegally places king in check
func moveAbandonsKing(board *ChessBoard, move Move, turn int8) bool {
	newBoard := doMove(move, *board)
	if turn == White {
		return newBoard.WhiteInCheck()
	} else {
		return newBoard.BlackInCheck()
	}
}
