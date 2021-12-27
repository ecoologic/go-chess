package main

import (
	"fmt"
	"testing"
)

func boardHas(board BoardType, origin positionType, expectedAtOrigin pieceType, destination positionType, expectedAtDestination pieceType) bool {
	return (board.at(origin) == expectedAtOrigin) && (board.at(destination) == expectedAtDestination)
}

func expectMove(t *testing.T, board BoardType, origin positionType, expectedAtOrigin pieceType, destination positionType, expectedAtDestination pieceType) {
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

	pieceByPosition := make(map[positionType]pieceType)
	pieceByPosition["B8"] = "N"
	pieceByPosition["F2"] = "p"
	pieceByPosition["A1"] = "r"
	pieceByPosition["D8"] = "Q"
	pieceByPosition["G1"] = "n"

	for position, piece := range pieceByPosition {
		if board.at(position) != piece {
			t.Error(fmt.Sprintf("board at %s should be a black knight, but was %s", position, piece))
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
