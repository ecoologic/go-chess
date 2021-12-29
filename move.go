package main

import (
	"fmt"
	"strings"
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
	pieceLetter := m.board.at(m.origin)
	// Both black and white
	switch strings.ToUpper(string(pieceLetter)) {
	case "P":
		return isPawnMoveLegal(m)
	default:
		return false
	}
}

func (m MoveType) deltaFile() int {
	return matrixColumnIndex(m.destination) - matrixColumnIndex(m.origin)
}

// TODO: include direction
func (m MoveType) deltaRank() int {
	return matrixRowIndex(m.destination) - matrixRowIndex(m.origin)
}

// Long: include destination
// TODO: incomplete logic, only implemented for pawn move by two
func (m MoveType) hasLongCorridor() bool {
	intermediatePosition := m.destination[0:1] + "3"
	return m.board.at(intermediatePosition) == "_" && m.board.at(m.destination) == "_"
}

func (m MoveType) originPieceLetter() pieceLetterType {
	return m.board.at(m.origin)
}

func (m MoveType) destinationPieceLetter() pieceLetterType {
	return m.board.at(m.destination)
}

func (m MoveType) isAttack() bool {
	originPiece := Piece{letter: m.originPieceLetter()}
	destinationPiece := Piece{letter: m.destinationPieceLetter()}

	return originPiece.isOpponent(destinationPiece)
}
