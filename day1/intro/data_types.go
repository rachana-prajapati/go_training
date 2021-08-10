package main

import "fmt"

//pkg level vars
var appname string = "My App"

func main() {
	//type1
	var msg string
	msg = "Hello World"
	fmt.Println(msg)

	//type2
	var msg2 string = "Hello2"
	fmt.Println(msg2)

	//type3 - type inference
	var msg3 = "Hello3"
	fmt.Println(msg3)

	//the following is applicable only in function. It can not be used at pkg level
	msg4 := "hello4"
	fmt.Println(msg4)

	/*
	--> option#0
	label = "Add Result: "
	x = 100
	y = 200
	*/

	/*
	--> option#1
	var label string
	var x,y int
	*/

	/*
	--> option#2
	var(
		label string
		x int
		y int
	)
	*/

	/*
	--> option#3
	var(
		label string
		x, y int
	)
	*/

	/*
	--> option#4
	var(
		label string = "Add Result: "
		x int = 100
		y int = 200
	)
	*/

	/*
	--> option#5
	var(
		label string = "Add Result: "
		x, y int = 100, 200
	)
	*/

	/*
	--> option#6 : type inference
	var(
		label, x, y = "Add Result: ", 100, 200
	)
	*/

	/*
	--> option#7
	var label, x, y = "Add Result: ", 100, 200
	*/

	/*
	--> option#8
	*/
	var label, x, y = "Add Result: ", 100, 200
	

	fmt.Println(label, x+y)
}