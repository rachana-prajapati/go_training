package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("Done")
}

func fibonacci(ch chan int) {
	x, y := 0, 1
	for i := 0; i < 20; i++ {
		ch <- x
		time.Sleep(100 * time.Millisecond)
		x, y = y, x+y
	}
	close(ch)
}
