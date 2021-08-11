package main

import "fmt"

func main() {
	c := make(chan int)

	go add(100, 200, c)

	result := <-c
	fmt.Println(result)
}

func add(x, y int, c chan int) {
	result := x + y
	c <- result
}
