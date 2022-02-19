package main

import (
	"log"
	"os"
)

func main() {
	for true {
		run()
	}
}

func run() {
	showBanner()

	fNum, err := askForFirstNumber()
	if err != nil {
		log.Println(err.Error())

		return
	}

	sNum, err := askForSecondNumber()
	if err != nil {
		log.Println(err.Error())

		return
	}

	op, err := askForOperation()
	if err != nil {
		log.Println(err.Error())

		return
	}

	if op < 1 || op > 4 {
		log.Println("Unavailable operation!")

		return
	}

	eq := Equation{fNum: fNum, sNum: sNum, op: op}

	calc(&eq)
	printResult(eq)
	continueOrExit()
}

func continueOrExit() {
	shouldContinueOrExit, err := askAboutContinueOrExit()
	if err != nil {
		log.Println(err.Error())

		return
	}

	switch shouldContinueOrExit {
	case 1:
		return
	case 2:
		os.Exit(2)
	default:
		continueOrExit()
	}
}
