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
	piece1 := boardAt(board.boardMatrix, "A1")

	if piece1 != "r" {
		t.Error(fmt.Sprintf("board at A1 should be a white rook, but was %s", piece1))
	}

	piece2 := boardAt(board.boardMatrix, "D8")
	if piece2 != "Q" {
		t.Error(fmt.Sprintf("board at D8 should be a black queen, but was %s", piece2))
	}

	piece3 := boardAt(board.boardMatrix, "G1")
	if piece3 != "n" {
		t.Error(fmt.Sprintf("board at G1 should be a white knight, but was %s", piece3))
	}
}

func TestBoardTypeAt(t *testing.T) {
	board := MakeBoard()

	piece1 := boardAt(board.boardMatrix, "B8")
	if piece1 != "N" {
		t.Error(fmt.Sprintf("board at B8 should be a black knight, but was %s", piece1))
	}
	piece2 := boardAt(board.boardMatrix, "F2")
	if piece2 != "p" {
		t.Error(fmt.Sprintf("board at F2 should be a white pawn, but was %s", piece2))
	}
}
