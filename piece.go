package main

type Piece struct {
	letter string
}

func (p Piece) isOpponent(otherPiece Piece) bool {
	return p.isWhite() != otherPiece.isWhite()
}

// A 65, Z 90 - a 97 z 122
func (p Piece) isWhite() bool {
	return int(p.letter[0]) < 90
}

// TODO: isPawn
