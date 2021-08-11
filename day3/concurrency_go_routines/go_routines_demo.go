package main

import (
	"fmt"
)

func main() {
	go Print("Hello")
	go Print("World")
	//time.Sleep(200 * time.Millisecond)
	var input string
	fmt.Scanln(&input)
}

func Print(s string) {
	fmt.Println(s)
}
