package myModule

import "fmt"

func addition(a, b int) int {
	return a + b
}

func DemoTheFunction() {
	firstNumber := 10
	secondNumber := 5
	hasil := addition(firstNumber, secondNumber)
	fmt.Println("firstNumber =", firstNumber, "secondNumber =", secondNumber)
	fmt.Println("Result  =", hasil)
}

func InputAdditionForOutsideFile(a, b int) {
	firstNumber := a
	secondNumber := b
	fmt.Println("--- InputAdditionForOutsideFile ---")
	fmt.Println("firstNumber =", firstNumber, "secondNumber =", secondNumber)
	fmt.Println("Result  =", addition(firstNumber, secondNumber))
}
