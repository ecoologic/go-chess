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
	move := MoveType{board, origin, destination, []MoveType{}}

	expectMoveIsLegal(t, move, false)
}

func TestRulePawnCanOpenByTwoWithAFreeCorridor(t *testing.T) {
	origin := "C2"
	destination := "C4"
	board := MakeBoard()

	legalMove := MoveType{board, origin, destination, []MoveType{}}
	expectMoveIsLegal(t, legalMove, true)

	board.Move("B1", "C3") // Occupy corridor

	illegalMove := MoveType{board, origin, destination, []MoveType{}}
	expectMoveIsLegal(t, illegalMove, false)
}

func TestRulePawnCanEatOnTheAngles(t *testing.T) {
	board := MakeBoard()
	board.Move("C7", "C3") // An opponent pawn to be eaten
	origin := "B2"
	destination := "C3"

	legalMove := MoveType{board, origin, destination, []MoveType{}}
	expectMoveIsLegal(t, legalMove, true)
}

func TestRulePawnCantEatBackwards(t *testing.T) {
	board := MakeBoard()
	board.Move("C7", "C3") // An opponent pawn to be eaten
	board.Move("B2", "D4") // Pawn positioned above at an angle
	origin := "D4"
	destination := "C3"

	legalMove := MoveType{board, origin, destination, []MoveType{}}
	expectMoveIsLegal(t, legalMove, false)
}

func TestRulePawnCanEatEnPassant(t *testing.T) {
	board := MakeBoard()
	previousMove := MoveType{board, "B2", "B4", []MoveType{}}
	board.Move("C7", "C4") // An opponent pawn on the attack
	board.Move("B2", "B4") // Pawn opens by two
	origin := "C4"
	destination := "B3"

	legalMove := MoveType{board, origin, destination, []MoveType{previousMove}}
	expectMoveIsLegal(t, legalMove, true)
}
