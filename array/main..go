package main

import (
	"fmt"
)

func main() {
	user_list := []string{"Alice", "Bob", "Charlie"}
	fmt.Println(user_list)

	special_list := [5]int{1: 10, 4: 20}
	fmt.Println(special_list)

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are"}
	fmt.Println(len(myslice2)) // Show Length of the Array
	fmt.Println(cap(myslice2)) // Show Capacity of the Array
	fmt.Println(myslice2)
	sliceExample()
	memoryEfficiency()
}

func sliceExample() {
	prices := []int{100, 200, 300, 400, 500}

	fmt.Println(prices[0])
	fmt.Println(prices[2])
}

func memoryEfficiency() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
  // Original slice
  fmt.Printf("numbers = %v\n", numbers)
  fmt.Printf("length = %d\n", len(numbers))
  fmt.Printf("capacity = %d\n", cap(numbers))

  // Create copy with only needed numbers
  neededNumbers := numbers[:len(numbers)-10]
  numbersCopy := make([]int, len(neededNumbers))
  copy(numbersCopy, neededNumbers)

  fmt.Printf("numbersCopy = %v\n", numbersCopy)
  fmt.Printf("length = %d\n", len(numbersCopy))
  fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
