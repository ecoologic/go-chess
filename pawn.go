package main

func validatePawnMove(move MoveType) bool {
	return (
	// Move by 1
	(move.deltaFile() == 0 && move.deltaRank() == 1 && move.AtDestination() == "_") ||
		// Move by 2 (opening)
		(move.deltaFile() == 0 && move.deltaRank() == 2 && move.hasLongCorridor()) ||
		// Eat
		(abs(move.deltaFile()) == 1 && move.deltaRank() == 1) && move.AtDestination() != "_")
	// TODO: opponent piece
}
