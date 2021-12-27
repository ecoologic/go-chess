package main

func validatePawnMove(move MoveType) bool {
	return ((move.deltaFile() == 0 && move.deltaRank() == 1 && move.AtDestination() == "_") ||
		(move.deltaFile() == 0 && move.deltaRank() == 2 && move.hasLongCorridor()))
}
