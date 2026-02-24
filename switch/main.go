package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday().String()

	switch day {
		case "Monday":
			fmt.Println("Start of the work week.")
		case "Tuesday":
			fmt.Println("Second day of the work week.")
		case "Wednesday":
			fmt.Println("Midweek day.")
		case "Thursday":
			fmt.Println("Almost the weekend.")
		case "Friday":
			fmt.Println("Last day of the work week.")
		case "Saturday", "Sunday":
			fmt.Println("It's the weekend!")
		default:
			fmt.Println("Unknown day.")
	}
}
