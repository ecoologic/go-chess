package main

import (
	"fmt"
	"strings"
)

type MoveType struct {
	board       BoardType
	origin      notationType
	destination notationType
}

func (m MoveType) String() string {
	return fmt.Sprintf("Move '%s' from %s to %s", m.board.at(m.origin), m.origin, m.destination)
}

func (m MoveType) isLegal() bool {
	piece := m.board.at(m.origin)
	// Both black and white
	switch strings.ToUpper(string(piece)) {
	case "P":
		return legalatePawnMove(m)
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
	intermediateNotation := m.destination[0:1] + "3"
	return m.board.at(intermediateNotation) == "_" && m.board.at(m.destination) == "_"
}

func (m MoveType) originPiece() pieceType {
	return m.board.at(m.origin)
}

func (m MoveType) destinationPiece() pieceType {
	return m.board.at(m.destination)
}

func (m MoveType) isAttack() bool {
	originRichPiece := RichPiece{letter: m.originPiece()}
	destinationRichPiece := RichPiece{letter: m.destinationPiece()}

	return originRichPiece.isOpponent(destinationRichPiece)
}

type RichPiece struct {
	letter string
}

func (rp RichPiece) isOpponent(otherPiece RichPiece) bool {
	return rp.isWhite() != otherPiece.isWhite()
}

// A 65, Z 90 - a 97 z 122
func (rp RichPiece) isWhite() bool {
	return int(rp.letter[0]) < 90
}

// TODO: isPawn
