package myModule

import "fmt"

type Student struct {
	Name     string
	Age      int
	Majoring string
}

func ShowStudentStruct() {
	mhs := Student{"Alice", 20, "Computer Science"}
	fmt.Printf("Name: %s, Age: %d, Majoring: %s\n", mhs.Name, mhs.Age, mhs.Majoring)
}

func ShowStudentStructWithNamedArgument() {
	student := Student{
		Name:     "Rudi",
		Majoring: "Biology",
		Age:      20,
	}
	fmt.Printf("Nama: %s, Umur: %d, Jurusan: %s\n", student.Name, student.Age, student.Majoring)
}
