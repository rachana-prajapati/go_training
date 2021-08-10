package main

import "fmt"

func main() {
	add := fn()
	res := add(100, 200)
	fmt.Printf("%d\n",res)
}

func fn() func(int, int) int{
	fmt.Printf("returning a func..\n")
	return func(x, y int) int{
		return x+y;
	}
}