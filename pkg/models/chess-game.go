package models

type ChessGame struct {
	currentState     ChessState
	currentMoveIndex int
	moves            []Move
}
