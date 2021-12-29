package main

import (
	"fmt"
	"testing"
)

func expectMoveIsValid(t *testing.T, move MoveType, expectIsValid bool) {
	if move.IsValid() == expectIsValid {
		return
	}
	t.Error(fmt.Sprintf("%v should be valid (%t) but it's the opposite", move.String(), expectIsValid))
}

func TestRulePawnCantMoveForwardWhenTheDestinationIsOccupied(t *testing.T) {
	origin := "C2"
	destination := "C3"
	board := MakeBoard()
	board.Move("B1", "C3") // Occupy destination
	move := MoveType{board, origin, destination}

	expectMoveIsValid(t, move, false)
}

func TestRulePawnCanOpenByTwoWithAFreeCorridor(t *testing.T) {
	origin := "C2"
	destination := "C4"
	board := MakeBoard()

	validMove := MoveType{board, origin, destination}
	expectMoveIsValid(t, validMove, true)

	board.Move("B1", "C3") // Occupy corridor

	invalidMove := MoveType{board, origin, destination}
	expectMoveIsValid(t, invalidMove, false)
}

// TODO: eat (direction, history for en-passant)

func TestRulePawnCanEatOnTheAngles(t *testing.T) {
	board := MakeBoard()
	board.Move("B7", "C3") // an opponent pawn to be eaten
	origin := "B2"
	destination := "C3"

	validMove := MoveType{board, origin, destination}
	expectMoveIsValid(t, validMove, true)
}
