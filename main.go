// https://cdn.britannica.com/71/7471-004-C94F7C98/chessmen-Position-beginning-game-queen-rook-king.jpg
package main

import (
	"fmt"
)

func main() {
	board := MakeBoard()
	fmt.Println("The board is: ", board)
	fmt.Println("D8 is: ", string(board.at("D8")))
}

/////////////

// type MoveType struct {
// 	board       BoardType
// 	origin      positionType
// 	destination positionType
// }

// func (move MoveType) Unchecked() {
// 	move.board.move(move.origin, move.destination)
// }
