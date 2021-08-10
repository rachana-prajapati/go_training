package main

import "fmt"

func main() {
	n1, n2 := 20, 3
	fmt.Println(add(n1, n2))
	fmt.Println(divide(n1, n2))
}

func add(n1 int, n2 int) int{
	return n1+n2
}

func divide(n1 int, n2 int) (quotient int, remainder int) {
	quotient = n1/n2
	remainder = n1%n2
	return quotient, remainder
}