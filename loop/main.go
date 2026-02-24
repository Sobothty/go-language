package main

import (
	"fmt"
)

func main() {
	for i := range 5 {
		fmt.Print("Number ", i, " ")
	}
	loopContinue()
	loopBreak()
	loopArray()
}

func loopContinue() {
	for i := range 5 {
		if i == 3 {
			continue
		}
		fmt.Print(i)
	}
}

func loopBreak() {
	for i := range 5 {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}

func loopArray() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}

	for i := range len(adj) {
		for j := range len(fruits) {
			fmt.Println(adj[i], fruits[j])
		}
	}
}
