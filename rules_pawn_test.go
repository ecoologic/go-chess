package main

import (
	"fmt"
	"testing"
)

func expectMoveIsLegal(t *testing.T, move MoveType, expectIsLegal bool) {
	if move.isLegal() == expectIsLegal {
		return
	}
	t.Error(fmt.Sprintf("%v should be legal (%t) but it's the opposite", move.String(), expectIsLegal))
}

func TestRulePawnCantMoveForwardWhenTheDestinationIsOccupied(t *testing.T) {
	origin := "C2"
	destination := "C3"
	board := MakeBoard()
	board.Move("B1", "C3") // Occupy destination
	move := MoveType{board, origin, destination}

	expectMoveIsLegal(t, move, false)
}

func TestRulePawnCanOpenByTwoWithAFreeCorridor(t *testing.T) {
	origin := "C2"
	destination := "C4"
	board := MakeBoard()

	legalMove := MoveType{board, origin, destination}
	expectMoveIsLegal(t, legalMove, true)

	board.Move("B1", "C3") // Occupy corridor

	illegalMove := MoveType{board, origin, destination}
	expectMoveIsLegal(t, illegalMove, false)
}

func TestRulePawnCanEatOnTheAngles(t *testing.T) {
	board := MakeBoard()
	board.Move("B7", "C3") // An opponent pawn to be eaten
	origin := "B2"
	destination := "C3"

	legalMove := MoveType{board, origin, destination}
	expectMoveIsLegal(t, legalMove, true)
}

func TestRulePawnCantEatBackwards(t *testing.T) {
	board := MakeBoard()
	board.Move("B7", "C3") // An opponent pawn to be eaten
	board.Move("B2", "D4") // Pawn positioned above at an angle
	origin := "D4"
	destination := "C3"

	legalMove := MoveType{board, origin, destination}
	expectMoveIsLegal(t, legalMove, false)
}

// TODO: eat (direction, history for en-passant)
