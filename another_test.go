package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 5 {
		t.Error("Expected 2 + 2 to equal 4")
	}

}
