package myModule

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person // Embedding Person struct
	Role   string
}

func (p *Person) PrintPersonDetails() {
	fmt.Println("Name :", p.Name)
	fmt.Println("Age :", p.Age)
}
