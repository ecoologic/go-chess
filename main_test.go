package main

import (
	"fmt"
	"testing"
)

func TestNewBoardMatrix(t *testing.T) {
	board := makeBoard()
	A8 := board[7][0]
	D1 := board[0][3]
	E5 := board[4][4]

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
