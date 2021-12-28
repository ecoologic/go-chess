// https://cdn.britannica.com/71/7471-004-C94F7C98/chessmen-Notation-beginning-game-queen-rook-king.jpg
package main

import (
	"fmt"
	"math"
)

func main() {
	board := MakeBoard()
	fmt.Println("The board is: ", board)
	fmt.Println("D8 is: ", string(board.at("D8")))
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

/////////////

// type MoveType struct {
// 	board       BoardType
// 	origin      notationType
// 	destination notationType
// }

// func (move MoveType) Unchecked() {
// 	move.board.move(move.origin, move.destination)
// }
