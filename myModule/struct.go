package myModule

import "fmt"

type Student struct {
	Name     string
	Age      int
	Majoring string
}

func ShowStudentStruct() {
	mhs := Student{"Alice", 20, "Informatics"}
	fmt.Printf("Name: %s, Age: %d, Majoring: %s\n", mhs.Name, mhs.Age, mhs.Majoring)
}
