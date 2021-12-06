package main

import (
	"fmt"
	"testing"
)

func TestNewBoardMatrix(t *testing.T) {
	board := MakeBoard()
	A8 := board.boardMatrix[7][0]
	D1 := board.boardMatrix[0][3]
	E5 := board.boardMatrix[4][4]

	if A8 != 'R' {
		t.Error(fmt.Sprintf("board should start with a black rook in A8, but was %c", A8))
	}
	if D1 != 'q' {
		t.Error(fmt.Sprintf("board should start with a white queen in D1, but was %c", D1))
	}
	if E5 != '_' {
		t.Error(fmt.Sprintf("board should start nothing in E5, but was %c", E5))
	}
}

func TestBoardAt(t *testing.T) {
	board := MakeBoard()

	pieceByPosition := make(map[positionType]pieceType)
	pieceByPosition["B8"] = 'N'
	pieceByPosition["F2"] = 'p'
	pieceByPosition["A1"] = 'r'
	pieceByPosition["D8"] = 'Q'
	pieceByPosition["G1"] = 'n'

	for position, piece := range pieceByPosition {
		if board.At(position) != piece {
			t.Error(fmt.Sprintf("board at %s should be a black knight, but was %c", position, piece))
		}
	}
}

// func TestMove(t *testing.T) {
// 	const board = MakeBoard()

// 	board = Move(board, "E2", "E3")
// }
