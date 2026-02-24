package main

import (
	"fmt"
	"time"
)

func main() {
	orders := make(chan string)

	go func() {
		fmt.Println("Waiter: Ready with order")
		orders <- "Steak, medium rare"
		fmt.Println("Waiter: Order sent to kitchen")
	}()

	go func() {
		fmt.Println("Chef: Waiting for orders...")
		order := <-orders
		fmt.Println("Chef: Received order:", order)
	}()

	time.Sleep(time.Second * 2)
}
