package main

import (
	"fmt"
)

type MoveType struct {
	board       BoardType
	origin      positionType
	destination positionType
}

func (m MoveType) String() string {
	return fmt.Sprintf("Move '%s' from %s to %s", m.board.at(m.origin), m.origin, m.destination)
}

func (m MoveType) isLegal() bool {
	piece := m.board.at(m.origin)
	// Both black and white
	switch {
	case piece.isPawn():
		return isPawnMoveLegal(m)
	default:
		return false
	}
}

func (m MoveType) deltaFile() int {
	return matrixColumnIndex(m.destination) - matrixColumnIndex(m.origin)
}

func (m MoveType) deltaRank() int {
	return m.originPiece().adjustDeltaByDirection(matrixRowIndex(m.destination) - matrixRowIndex(m.origin))
}

// Long: include destination
// TODO: incomplete logic, only implemented for pawn move by two
func (m MoveType) hasLongCorridor() bool {
	intermediatePosition := m.destination[0:1] + "3"
	return m.board.at(intermediatePosition).isNil() && m.board.at(m.destination).isNil()
}

func (m MoveType) originPiece() Piece {
	return m.board.at(m.origin)
}

func (m MoveType) destinationPiece() Piece {
	return m.board.at(m.destination)
}

func (m MoveType) isAttack() bool {
	return m.originPiece().isOpponent(m.destinationPiece())
}

func (m MoveType) pawnWasPassing() bool {
	return false
}
