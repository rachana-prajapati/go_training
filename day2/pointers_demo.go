package main

import "fmt"

func main() {
	x := 10
	var xPtr *int = &x
	fmt.Println("x -> ", x)
	fmt.Println("xPtr -> ", xPtr)

	//dereferencing the pointer
	y := *xPtr
	fmt.Println("y -> ", y)

	fmt.Println("before inc, x -> ", x)
	inc(&x)
	fmt.Println("after inc, x -> ", x)

	n1, n2 := 10, 20
	fmt.Println("before swapping -> ", n1, n2)
	swap(&n1, &n2)
	fmt.Println("after swapping -> ", n1, n2)
}

func inc(xPtr *int) {
	*xPtr++
}

func swap(n1Ptr, n2Ptr *int) {
	*n1Ptr, *n2Ptr = *n2Ptr, *n1Ptr
}
