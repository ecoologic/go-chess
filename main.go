package main

import (
	"fmt"
	"strconv"
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

// file: letter
// rank: number

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

func positionToCoordinates(position string) (int, int) {
	fileCoord := 0
	rankCoord, _ := strconv.Atoi(position[1:1])
	return fileCoord, rankCoord
}

func boardAt(board boardMatrixType, position string) string {
	fileCoord, rankCoord := positionToCoordinates(position)
	return board[rankCoord][fileCoord]
}

func main() {
	board := boardType{}
	fmt.Println("The board is: ", board.boardMatrix)
	fmt.Println("A1 is: ", boardAt(makeBoard(), "A1"))
}

// func main() {
// 	fmt.Println("Hello World", Calculate(5))
// 	var x int = 3
// 	var a [4]int
// 	b := []int{5, 4, 3, 2, 1}
// 	b = append(b, 0)
// 	c := make(map[string]int)
// 	c["ok"] = 1
// 	c["sure"] = 2
// 	delete(c, "sure")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}
// 	for i, v := range b {
// 		fmt.Println(i, sum(v, 10))
// 	}

// 	p := person{name: "erik", age: 41}
// 	fmt.Println(x, a, b, c, p)

// 	// pointers

// 	n := 7
// 	inc(&n)
// 	fmt.Println("pointers: ", n)
// }

// func inc(n *int) {
// 	*n++
// }

// type person struct {
// 	name string
// 	age  int
// }

// // Calculate returns x + 2.
// func Calculate(x int) (result int) {
// 	result = x + 2
// 	return result
// }

// func sum(x int, y int) int {
// 	return x + y
// }
