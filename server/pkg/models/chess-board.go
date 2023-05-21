package models

import (
	"fmt"
)

// 7	BR BN BB BQ BK BB BN BR	|
// 6	BP BP BP BP BP BP BP BP	|
// 5	-- -- -- -- -- -- -- --	|
// 4	-- -- -- -- -- -- -- --	|
// 3	-- -- -- -- -- -- -- --	|
// 2	-- -- -- -- -- -- -- --	|
// 1	WP WP WP WP WP WP WP WP	|
// 0	WR WN WB WQ WK WB WN WR	|
// i
//	j   0  1  2  3  4  5  6  7

type ChessBoard [8][8]int8

func NewChessBoard() *ChessBoard {
	return &ChessBoard{
		{WhiteRook, WhiteKnight, WhiteBishop, WhiteQueen, WhiteKing, WhiteBishop, WhiteKnight, WhiteRook},
		{WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn, WhitePawn},
		{EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare},
		{EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare},
		{EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare},
		{EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare, EmptySquare},
		{BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn, BlackPawn},
		{BlackRook, BlackKnight, BlackBishop, BlackQueen, BlackKing, BlackBishop, BlackKnight, BlackRook},
	}
}

func (board *ChessBoard) printChessBoard() {
	for i := 7; i >= 0; i-- {
		fmt.Println("-----------------------------------------------------------------")
		for j := 0; j < 8; j++ {
			fmt.Printf("| %v\t", board[i][j])
		}
		fmt.Println("|")
	}
	fmt.Println("-----------------------------------------------------------------")
}

func (board *ChessBoard) IsSquareAttackedByWhite(i, j int) bool {
	// check for pawn attacks
	if i-1 >= 0 && j-1 >= 0 && board[i-1][j-1] == WhitePawn {
		return true
	}
	if i-1 >= 0 && j+1 <= 7 && board[i-1][j+1] == WhitePawn {
		return true
	}
	// check for knight attacks
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
	// check for bishop/queen attacks
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

	// check for rook/queen attacks
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

func (board *ChessBoard) IsSquareAttackedByBlack(i, j int) bool {
	// check for pawn attacks
	if i+1 <= 7 && j-1 >= 0 && board[i+1][j-1] == BlackPawn {
		return true
	}
	if i+1 <= 7 && j+1 <= 7 && board[i+1][j+1] == BlackPawn {
		return true
	}
	// check for knight attacks
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
	// check for bishop/queen attacks
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

	// check for rook/queen attacks
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

func (board *ChessBoard) IsWhiteInCheck() bool {
	// find king
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			if board[i][j] == WhiteKing {
				// check if square is attacked by
				return board.IsSquareAttackedByBlack(i, j)
			}
		}
	}
	// huh
	return false
}

func (board *ChessBoard) IsBlackInCheck() bool {
	// find king
	for j := 7; j >= 0; j-- {
		for i := 7; i >= 0; i-- {
			if board[i][j] == BlackKing {
				return board.IsSquareAttackedByWhite(i, j)
			}
		}
	}
	// huh
	return false
}

func executeMoveOnBoard(move Move, board ChessBoard) (ChessBoard, bool, bool, bool, bool) {

	whiteCanCastleShort := true
	whiteCanCastleLong := true
	blackCanCastleShort := true
	blackCanCastleLong := true

	if move.Type == CastleShort {
		row := move.OldSquare.Row

		// move king
		board[6][row] = board[4][row]

		// move rook
		board[5][row] = board[7][row]

		// clear spaces
		board[4][row] = EmptySquare
		board[7][row] = EmptySquare

		if row == 0 {
			whiteCanCastleShort = false
			whiteCanCastleLong = false
		} else if row == 7 {
			blackCanCastleShort = false
			blackCanCastleLong = false
		} else {
			fmt.Printf("Wait a minute... castling on rank %v???\n", row)
		}

		return board, whiteCanCastleShort, whiteCanCastleLong, blackCanCastleShort, blackCanCastleLong
	}
	if move.Type == CastleLong {
		row := move.OldSquare.Row

		// move king
		board[2][row] = board[4][row]

		// move rook
		board[3][row] = board[0][row]

		// clear spaces
		board[0][row] = EmptySquare
		board[1][row] = EmptySquare
		board[4][row] = EmptySquare

		if row == 0 {
			whiteCanCastleShort = false
			whiteCanCastleLong = false
		} else if row == 7 {
			blackCanCastleShort = false
			blackCanCastleLong = false
		} else {
			fmt.Printf("Wait a minute... castling on rank %v???\n", row)
		}

		return board, whiteCanCastleShort, whiteCanCastleLong, blackCanCastleShort, blackCanCastleLong
	}
	if move.Type == EnPassant {
		// move pawn
		board[move.NewSquare.Row][move.NewSquare.Col] = board[move.OldSquare.Row][move.OldSquare.Col]

		// clear spaces
		board[move.OldSquare.Row][move.OldSquare.Col] = EmptySquare
		board[move.OldSquare.Row][move.NewSquare.Col] = EmptySquare

		return board, whiteCanCastleShort, whiteCanCastleLong, blackCanCastleShort, blackCanCastleLong
	}

	movingPieceType := board[move.OldSquare.Row][move.OldSquare.Col]
	newRow := move.NewSquare.Row
	if move.Type == PromoteQueen {
		if newRow == 0 {
			movingPieceType = BlackQueen
		} else if newRow == 7 {
			movingPieceType = WhiteQueen
		} else {
			fmt.Printf("Wait a minute... promoting on rank %v???", newRow)
		}
	} else if move.Type == PromoteRook {
		if newRow == 0 {
			movingPieceType = BlackRook
		} else if newRow == 7 {
			movingPieceType = WhiteRook
		} else {
			fmt.Printf("Wait a minute... promoting on rank %v???", newRow)
		}
	} else if move.Type == PromoteBishop {
		if newRow == 0 {
			movingPieceType = BlackBishop
		} else if newRow == 7 {
			movingPieceType = WhiteBishop
		} else {
			fmt.Printf("Wait a minute... promoting on rank %v???", newRow)
		}
	} else if move.Type == PromoteKnight {
		if newRow == 0 {
			movingPieceType = BlackKnight
		} else if newRow == 7 {
			movingPieceType = WhiteKnight
		} else {
			fmt.Printf("Wait a minute... promoting on rank %v???", newRow)
		}
	}

	// make checks to see if move effects castling rights
	if movingPieceType == WhiteKing {
		whiteCanCastleShort = false
		whiteCanCastleLong = false
	} else if movingPieceType == BlackKing {
		blackCanCastleShort = false
		blackCanCastleLong = false
	} else if (move.OldSquare.Row == 0 && move.OldSquare.Col == 0) || (move.NewSquare.Row == 0 && move.NewSquare.Col == 0) {
		whiteCanCastleLong = false
	} else if (move.OldSquare.Row == 0 && move.OldSquare.Col == 7) || (move.NewSquare.Row == 0 && move.NewSquare.Col == 7) {
		whiteCanCastleShort = false
	} else if (move.OldSquare.Row == 7 && move.OldSquare.Col == 0) || (move.NewSquare.Row == 7 && move.NewSquare.Col == 0) {
		blackCanCastleLong = false
	} else if (move.OldSquare.Row == 7 && move.OldSquare.Col == 7) || (move.NewSquare.Row == 7 && move.NewSquare.Col == 7) {
		blackCanCastleShort = false
	}
	// make move
	board[move.NewSquare.Row][move.NewSquare.Col] = movingPieceType
	board[move.OldSquare.Row][move.OldSquare.Col] = EmptySquare

	return board, whiteCanCastleShort, whiteCanCastleLong, blackCanCastleShort, blackCanCastleLong
}
