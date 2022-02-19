package main

import "testing"

func TestCalcShouldAddNumbers(t *testing.T) {
	eq := Equation{fNum: 1, sNum: 2, op: 1}

	calc(&eq)
	if eq.r != 3 {
		t.Fail()
	}
}
func TestCalcShouldSubtractNumbers(t *testing.T) {
	eq := Equation{fNum: 1, sNum: 2, op: 2}

	calc(&eq)
	if eq.r != -1 {
		t.Fail()
	}
}
func TestCalcShouldMultiplyNumbers(t *testing.T) {
	eq := Equation{fNum: 2, sNum: 5, op: 3}

	calc(&eq)
	if eq.r != 10 {
		t.Fail()
	}
}
func TestCalcShouldDivideNumbers(t *testing.T) {
	eq := Equation{fNum: 10, sNum: 2, op: 4}

	calc(&eq)
	if eq.r != 5 {
		t.Fail()
	}
}
