package main

import (
	"fmt"
	"testing"
)

func boardHas(board BoardType, origin notationType, expectedAtOrigin pieceType, destination notationType, expectedAtDestination pieceType) bool {
	return (board.at(origin) == expectedAtOrigin) && (board.at(destination) == expectedAtDestination)
}

func expectMove(t *testing.T, board BoardType, origin notationType, expectedAtOrigin pieceType, destination notationType, expectedAtDestination pieceType) {
	if !boardHas(board, origin, expectedAtOrigin, destination, expectedAtDestination) {
		t.Error(fmt.Sprintf(
			"Expected move:\n%s == %s (actual %s)\n%s == %s (actual %s)",
			origin, expectedAtOrigin, board.at(origin),
			destination, expectedAtDestination, board.at(destination)))
	}

}

func TestNewBoardMatrix(t *testing.T) {
	board := MakeBoard()
	A8 := board.boardMatrix[7][0]
	D1 := board.boardMatrix[0][3]
	E5 := board.boardMatrix[4][4]

	if A8 != "R" {
		t.Error(fmt.Sprintf("board should start with a black rook in A8, but was %s", A8))
	}
	if D1 != "q" {
		t.Error(fmt.Sprintf("board should start with a white queen in D1, but was %s", D1))
	}
	if E5 != "_" {
		t.Error(fmt.Sprintf("board should start nothing in E5, but was %s", E5))
	}
}

func TestBoardAt(t *testing.T) {
	board := MakeBoard()

	pieceByNotation := make(map[notationType]pieceType)
	pieceByNotation["B8"] = "N"
	pieceByNotation["F2"] = "p"
	pieceByNotation["A1"] = "r"
	pieceByNotation["D8"] = "Q"
	pieceByNotation["G1"] = "n"

	for notation, piece := range pieceByNotation {
		if board.at(notation) != piece {
			t.Error(fmt.Sprintf("board at %s should be a black knight, but was %s", notation, piece))
		}
	}
}

func TestBoardMovePawnAheadValid(t *testing.T) {
	origin := "E2"
	destination := "E3"
	board := MakeBoard()

	board.Move(origin, destination)

	expectMove(t, board, origin, "_", destination, "p")
}
