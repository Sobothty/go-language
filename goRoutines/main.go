package main

import (
	"fmt"
)

func myMessage() {
	fmt.Println("I just got executed!")
}

func main(){
	go myMessage()
}
