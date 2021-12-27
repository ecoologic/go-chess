package main

import (
	"fmt"
	"testing"
)

func expectMoveIsValid(t *testing.T, move MoveType, expectIsValid bool) {
	if move.IsValid() == expectIsValid {
		return
	}
	t.Error(fmt.Sprintf("Move should be valid but it's not"))
}

func TestRulePawnCantMoveForwardWhenTheDestinationIsOccupied(t *testing.T) {
	origin := "E2"
	destination := "E3"
	board := MakeBoard()
	board.Move("B1", "C3") // Occupy destination
	move := MoveType{board, origin, destination}

	expectMoveIsValid(t, move, false)
}
