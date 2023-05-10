package models

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
// i							|
//
//	j   0  1  2  3  4  5  6  7	|
//
// ----------------------------+
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
