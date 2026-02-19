package main

import (
	"fmt"
)

func main() {
	var a bool = true
	var b int = 5
	var c float32 = 3.14
	var d string = "Hello World"
	var x float64 = 1.7e+308
	
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
