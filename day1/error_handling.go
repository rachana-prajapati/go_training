package main

import (
	"fmt"
	"errors"
)

var divideByZeroError = errors.New("divide by zero error")

func main() {
	fmt.Println(divide(10,2))
	fmt.Println(divide(10,0))

	res, err := divide(10, 0)
	if(err == divideByZeroError){
		fmt.Println("divisor can not be null.")
	} else if err != nil {
		fmt.Println("unknown error..")
	} else {
		fmt.Printf("res = %d", res)
	}
}

func divide (x, y int) (int, error) {
	if y != 0{
		return x/y, nil
	} else {
		return 0, divideByZeroError
	}
}