package main

import (
	"fmt"
	"time"
)

func main() {
	//var printString string
	var tunnel1 = make(chan int)
	var tunnel2 = make(chan int)
	str1, str2 := "Hello", "World"
	go Print(str1, 1*time.Second, tunnel1, tunnel2)
	go Print(str2, 2*time.Second, tunnel2, tunnel1)
	tunnel1 <- 1
	var input string
	fmt.Scanln(&input)
}

func Print(str string, delay time.Duration, tunnel1, tunnel2 chan int) {
	for i := 0; i < 5; i++ {
		<-tunnel1
		fmt.Println(str)
		time.Sleep(delay)
		tunnel2 <- 1
	}
}

/*
Output:
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
