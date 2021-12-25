package main

func validatePawnMove(move MoveType) bool {
	switch {
	case move.deltaFile() == 1 && move.AtDestination() != '_':
		return true
	default:
		return false
	}
}
