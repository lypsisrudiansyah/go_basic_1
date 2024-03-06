package myModule

import "fmt"

// func
type Animal interface {
	Speak() string
	Eat() string
}

type Lion struct {
	Name string
}

type Cow struct {
	Name string
}

func (l Lion) Speak() string {
	// return "Roar"
	return fmt.Sprintf("%s sounds Roar", l.Name)
}

func (l Lion) Eat() string {
	// return "Meat"
	return fmt.Sprintf("%s eats Meat", l.Name)
}

func (c Cow) Speak() string {
	return "Moo"
}

func (c Cow) Eat() string {
	return "Grass"
}
