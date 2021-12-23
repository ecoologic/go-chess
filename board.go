package main

import (
	"strings"
)

// BLACK uppercase
// X A B C D E F G H
// 8 R N B Q K B N R
// 7 P P P P P P P P
// 6 _ _ _ _ _ _ _ _
// 5 _ _ _ _ _ _ _ _
// 4 _ _ _ _ _ _ _ _
// 3 _ _ _ _ _ _ _ _
// 2 p p p p p p p p
// 1 r n b q k b n r
// X A B C D E F G H
// white lowercase

func MakeBoard() BoardType {
	board := boardMatrixType{
		{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}, // [0][0] (1A) White
		{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'},
		{'_', '_', '_', '_', '_', '_', '_', '_'},
		{'_', '_', '_', '_', '_', '_', '_', '_'},
		{'_', '_', '_', '_', '_', '_', '_', '_'}, // Assigned inverted
		{'_', '_', '_', '_', '_', '_', '_', '_'},
		{'P', 'P', 'P', 'P', 'P', 'P', 'P', 'P'},
		{'R', 'N', 'B', 'Q', 'K', 'B', 'N', 'R'}, // [7][7] (8H) Black
	}
	return BoardType{board}
}

//////////

type BoardType struct {
	boardMatrix boardMatrixType
}

func (b BoardType) at(position positionType) pieceType {
	return b.boardMatrix[matrixRowIndex(position)][matrixColumnIndex(position)]
}

func (b *BoardType) Move(origin positionType, destination positionType) {
	b.boardMatrix[matrixRowIndex(destination)][matrixColumnIndex(destination)] =
		b.boardMatrix[matrixRowIndex(origin)][matrixColumnIndex(origin)]
}

func validatePawnMove(move MoveType) bool {
	switch {
	case move.deltaFile() == 1 && move.AtDestination() != '_':
		return true
	default:
		return false
	}
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
