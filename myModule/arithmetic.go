package myModule

import "fmt"

func PrintArithmetic() {
	var firstNumber int = 10
	var secondNumber int = 5

	fmt.Println("==== Arithmetic Operations ====")
	fmt.Println("First Number :", firstNumber)
	fmt.Println("Second Number :", secondNumber)
	fmt.Println("Addition :", firstNumber+secondNumber)
	fmt.Println("Subtraction :", firstNumber-secondNumber)
	fmt.Println("Multiplication :", firstNumber*secondNumber)
	fmt.Println("Division :", firstNumber/secondNumber)
	fmt.Println("Modulus :", firstNumber%secondNumber)
}
