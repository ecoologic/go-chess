// https://cdn.britannica.com/71/7471-004-C94F7C98/chessmen-Position-beginning-game-queen-rook-king.jpg
package ecoologicChess

import (
	"fmt"
	"strconv"
	"strings"
)

//////////

// squares
// ??? position -> notation
type positionType = string
type pieceType rune
type boardMatrixType [8][8]pieceType

// board | notation | index
// file  | letter   | column
// rank  | number   | row
func matrixColumnIndex(position positionType) (result int) {
	result = int(strings.ToUpper(position)[0]) - int('A')
	return // naked return
}

func matrixRowIndex(position positionType) int {
	rank1to8, _ := strconv.Atoi(string(position[1]))
	return rank1to8 - 1
}

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
