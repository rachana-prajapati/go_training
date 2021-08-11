package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}
var opCount = 0
var mutex = &sync.Mutex{}

func main() {
	wg.Add(1)
	go Add(10, 20)

	wg.Add(1)
	go Add(100, 200)

	wg.Wait()
	fmt.Println("opCount -> ", opCount)
	fmt.Println("Exiting the main fn")
}

func Add(x, y int) {
	res := x + y
	fmt.Println(x, y, " Add -> ", res)
	mutex.Lock()
	{
		opCount++
	}
	mutex.Unlock()
	wg.Done()
}
