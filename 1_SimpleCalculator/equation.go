package main

type Equation struct {
	fNum, sNum, r, op int
}

func calc(eq *Equation) {
	switch eq.op {
	case 1:
		add(eq)
	case 2:
		subtract(eq)
	case 3:
		multiply(eq)
	case 4:
		divide(eq)
	}
}

func add(eq *Equation) {
	eq.r = eq.fNum + eq.sNum
}

func subtract(eq *Equation) {
	eq.r = eq.fNum - eq.sNum
}

func multiply(eq *Equation) {
	eq.r = eq.fNum * eq.sNum
}

func divide(eq *Equation) {
	eq.r = eq.fNum / eq.sNum
}
