package main

import "strings"

type Piece struct {
	letter string
}

func (p Piece) isOpponent(otherPiece Piece) bool {
	if p.isNil() || otherPiece.isNil() {
		return false
	} else {
		return p.isWhite() != otherPiece.isWhite()
	}
}

// A 65, Z 90 - a 97 z 122
func (p Piece) isWhite() bool {
	return int(p.letter[0]) < 90
}

func (p Piece) isNil() bool {
	return p.letter == "_"
}

func (p Piece) isPawn() bool {
	return strings.ToUpper(string(p.letter)) == "P"
}

func (p Piece) adjustDeltaByDirection(delta int) int {
	if p.isWhite() {
		return 0 - delta
	} else {
		return delta
	}
}
