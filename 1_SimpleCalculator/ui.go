package main

import (
	"errors"
	"fmt"
)

func showBanner() {
	println("### CALCULATOR ###")
}

func askForFirstNumber() (int, error) {
	print("Enter first number: ")
	return getNumber()
}

func askForSecondNumber() (int, error) {
	print("Enter second number: ")
	return getNumber()
}

func askForOperation() (int, error) {
	println("What should I do with these numbers? ")
	println("1. Add")
	println("2. Subtract")
	println("3. Multiply")
	println("4. Divide")
	print("Choose an operation number (1-4): ")

	return getNumber()
}

func getNumber() (int, error) {
	var num int
	scan, err := fmt.Scan(&num)

	if err != nil {
		return 0, err
	}

	if scan == 0 {
		return 0, errors.New("you need to provide a number")
	}

	return num, nil
}

func printResult(eq Equation) {
	println("Result:", eq.r)
}

func askAboutContinueOrExit() (int, error) {
	println("Continue (1) | Exit (2)")

	return getNumber()
}
