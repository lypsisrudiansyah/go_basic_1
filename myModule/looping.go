package myModule

import "fmt"

func DoTheLooping() {
	// * Simple looping
	fmt.Println("---- Looping 1 ----")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// * Looping with array
	theNumber := []int{1, 2, 3, 4, 5}
	fmt.Println("---- Looping 2 ----")
	for i := 0; i < len(theNumber); i++ {
		fmt.Println("index", i, " - value : ", theNumber[i])
	}

	fmt.Println("---- Looping 3 ----")
	// * Look like while loop
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
}
