package models

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

// masks
const (
	TurnMask    = 0b0000000000000001
	StartIMask  = 0b0000000000001110
	StartJMask  = 0b0000000001110000
	EndIMask    = 0b0000001110000000
	EndJMask    = 0b0001110000000000
	SpecialMask = 0b1110000000000000
)

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

func decodeMove(move Move) (turn, specialMove, i1, j1, i2, j2 uint16) {
	encoded := uint16(move)
	turn = encoded & TurnMask
	i1 = (encoded & StartIMask) >> 1
	j1 = (encoded & StartJMask) >> 4
	i2 = (encoded & EndIMask) >> 7
	j2 = (encoded & EndJMask) >> 10
	specialMove = (encoded & SpecialMask) >> 13
	return turn, specialMove, i1, j1, i2, j2
}
