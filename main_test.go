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
		if board.at(position) != piece {
			t.Error(fmt.Sprintf("board at %s should be a black knight, but was %c", position, piece))
		}
	}
}

func TestBoardMovePawnAheadValid(t *testing.T) {
	origin := "E2"
	destination := "E3"
	board := MakeBoard()
	piece := board.at(origin)

	board.Move(origin, destination)

	if board.at(origin) != '_' && board.at(destination) != piece {
		t.Error(fmt.Sprintf(
			"Move Pawn ahead valid error: E2 should be '_' and E3 'p', but was '%c' and '%c'",
			board.at(origin), board.at(destination)), board)
	}
}
func TestBoardMovePawnAheadOccupied(t *testing.T) {
	origin := "E2"
	destination := "E3"
	board := MakeBoard()
	piece := board.at(origin)
	board.Move("B1", "C3") // Occupy destination

	board.Move(origin, destination)

	if board.at(origin) != piece && board.at(destination) != '_' {
		t.Error(fmt.Sprintf(
			"Move Pawn ahead occupied error: origin '%c' and destination '%c'",
			board.at(origin), board.at(destination)), board)
	}
}
