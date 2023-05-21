package models

type MoveType int

const (
	Normal MoveType = iota
	CastleShort
	CastleLong
	EnPassant
	PromoteQueen
	PromoteRook
	PromoteBishop
	PromoteKnight
)

type Location struct {
	Row int
	Col int
}

type Move struct {
	Type      MoveType
	OldSquare Location
	NewSquare Location
}

func NewMove(moveType MoveType, oldI, oldJ, newI, newJ int) Move {
	return Move{
		Type: moveType,
		OldSquare: Location{
			Row: oldI,
			Col: oldJ,
		},
		NewSquare: Location{
			Row: newI,
			Col: newJ,
		},
	}
}
