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
	Row int `json:"row"`
	Col int `json:"col"`
}

type Move struct {
	Type      MoveType `json:"type"`
	OldSquare Location `json:"oldSquare"`
	NewSquare Location `json:"newSquare"`
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
