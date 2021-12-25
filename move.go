package main

import "strings"

func (b *BoardType) Move(origin positionType, destination positionType) {
	b.boardMatrix[matrixRowIndex(destination)][matrixColumnIndex(destination)] =
		b.boardMatrix[matrixRowIndex(origin)][matrixColumnIndex(origin)]
}

//////////

type MoveType struct {
	board       BoardType
	origin      positionType
	destination positionType
}

func (m MoveType) IsValid() bool {
	piece := m.board.at(m.origin)
	// Both black and white
	switch strings.ToUpper(string(piece)) {
	case "P":
		return validatePawnMove(m)
	default:
		return false
	}
}

func (m MoveType) deltaFile() int {
	return matrixColumnIndex(m.destination) - matrixColumnIndex(m.origin)
}

func (m MoveType) AtDestination() pieceType {
	return m.board.at(m.destination)
}
