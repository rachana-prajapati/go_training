package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = "string"
	x = true
	x = []int{1, 2, 3}
	fmt.Println("x -> ", x)

	var val interface{}
	val = 100
	if number, ok := val.(int); ok {
		res := number + 200
		fmt.Println("res -> ", res)
	} else {
		fmt.Println("not of type int")
	}

	switch v := val.(type) {
	case int:
		fmt.Println(v, " is of type int")
	case string:
		fmt.Println(v, " is of type string")
	case []int:
		fmt.Println(v, " is of type slice")
	default:
		fmt.Println("unknown type")
	}
}
