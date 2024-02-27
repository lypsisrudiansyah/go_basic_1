package myModule

import "fmt"

// * Example with set the length of the array
func TheArray1() {
	var angka [5]int
	angka[0] = 1
	angka[1] = 2
	angka[2] = 3
	angka[3] = 4
	angka[4] = 5

	fmt.Println(angka)
}

// * Without set the length of the array
func TheArray2() {
	angka := []int{1, 2, 4, 4, 5, 7}

	fmt.Println(angka)
}
