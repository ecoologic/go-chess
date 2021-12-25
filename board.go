package main

import (
	"strconv"
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

//////////

// squares
// ??? position -> notation
type positionType = string
type pieceType rune
type boardMatrixType [8][8]pieceType

// board | notation | index
// file  | letter   | column
// rank  | number   | row
func matrixColumnIndex(position positionType) (result int) {
	result = int(strings.ToUpper(position)[0]) - int('A')
	return // naked return
}

func matrixRowIndex(position positionType) int {
	rank1to8, _ := strconv.Atoi(string(position[1]))
	return rank1to8 - 1
}
