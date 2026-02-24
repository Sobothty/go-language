package main

import (
	"fmt"
)

func main() {
	var a = make(map[string]string)
	a["name"] = "Sobothty"
	a["age"] = "20"
	a["gender"] = "Male"

	fmt.Printf("Name: %s\n", a["name"])
	fmt.Printf("Age: %s\n", a["age"])
	fmt.Printf("Gender: %s\n", a["gender"])

	b := map[string]string{
		"name":   "Sobothty",
		"age":    "20",
		"gender": "Male",
	}

	fmt.Printf("Name: %s\n", b["name"])
	fmt.Printf("Age: %s\n", b["age"])
	fmt.Printf("Gender: %s\n", b["gender"])
}
