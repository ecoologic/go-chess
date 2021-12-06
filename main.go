package main

import (
	"fmt"
	"strconv"
	"strings"
)

type boardMatrixType [8][8]string

type boardType struct {
	boardMatrix boardMatrixType
}

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

func makeBoard() boardMatrixType {
	return boardMatrixType{
		{"r", "n", "b", "q", "k", "b", "n", "r"}, // [0][0] (1A) White
		{"p", "p", "p", "p", "p", "p", "p", "p"},
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_", "_", "_", "_"}, // Assigned inverted
		{"_", "_", "_", "_", "_", "_", "_", "_"},
		{"P", "P", "P", "P", "P", "P", "P", "P"},
		{"R", "N", "B", "Q", "K", "B", "N", "R"}, // [7][7] (8H) Black
	}
}

// board, position, coord, index
// file, letter, y, column
// rank, number, x, row
func matrixColumnIndex(position string) (result int) {
	result = int(strings.ToUpper(position)[0]) - int("A"[0])
	return // naked return
}

func matrixRowIndex(position string) int {
	rank1to8, _ := strconv.Atoi(string(position[1]))
	return rank1to8 - 1
}

func boardAt(board boardMatrixType, position string) string {
	return board[matrixRowIndex(position)][matrixColumnIndex(position)]
}

func main() {
	board := makeBoard()
	fmt.Println("The board is: ", board)
	fmt.Println("A1 is: ", boardAt(board, "A1"))
}
