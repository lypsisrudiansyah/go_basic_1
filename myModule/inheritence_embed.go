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

func PrintPersonDetails() {
	employeeData := Employee{
		Person: Person{
			Name: "Rudiansyah", Age: 22,
		},
		Role: "Software Engineer",
	}

	fmt.Printf("Name : %s, Age : %d, Role : %s \n", employeeData.Name, employeeData.Age, employeeData.Role)
}
