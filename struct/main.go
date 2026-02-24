package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	var person1 Person

	person1.Name = "Sobothty"
	person1.Age = 20
	person1.Gender = "Male"

	fmt.Printf("Name: %s\n", person1.Name)
	fmt.Printf("Age: %d\n", person1.Age)
	fmt.Printf("Gender: %s\n", person1.Gender)

	printPersonInfo(person1)
}

func printPersonInfo(p Person) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Age: %d\n", p.Age)
	fmt.Printf("Gender: %s\n", p.Gender)
}
