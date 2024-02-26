package myModule

import "fmt"

func PrintUserDetails() {
	var age int = 25
	age = 26

	var height float64 = 165.2

	var name string = ""
	name = "John Doe"

	var active bool = false
	active = true

	fmt.Println("Age :", age)
	fmt.Println("Height :", height)
	fmt.Println("Name :", name)
	fmt.Println("Status(Active) :", active)
}
