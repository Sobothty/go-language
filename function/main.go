package main

import (
	"fmt"
)

func main() {
	myMessage()
	myName("Sobothty")
	result := returnWithType(5, 8)
	fmt.Printf("The sum is: %d\n", result)

	returnWithVal(5, 9)
}

func myMessage() {
	fmt.Println("I just got executed!")
}

func myName(name string) {
	fmt.Printf("My name is %s\n", name)
}

func returnWithType(a int, b int) int {
	return a + b
}

func returnWithVal(x int, y int) (result int) {
  result = x + y
  fmt.Printf("The sum is: %d\n", result)
  return 
}