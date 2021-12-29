package main

func legalatePawnMove(move MoveType) bool {
	// fmt.Println(fmt.Sprintf(">>> pawn legalatePawnMove"),
	// 	move.deltaFile(), move.deltaRank(), move.board.at(move.origin), move.destinationPiece(), move.isAttack())
	return (
	// Move by 1
	(move.deltaFile() == 0 && move.deltaRank() == 1 && move.destinationPiece() == "_") ||
		// Move by 2 (opening)
		(move.deltaFile() == 0 && move.deltaRank() == 2 && move.hasLongCorridor()) ||
		// Eat
		(abs(move.deltaFile()) == 1 && move.deltaRank() == 1) && move.isAttack())
}
