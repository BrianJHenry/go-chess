package models

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
