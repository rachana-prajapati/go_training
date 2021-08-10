package main

import "fmt"

func main() {

	for {
		var input int
		fmt.Println("-----------------");
		fmt.Println("----Main menu----");
		fmt.Println("-----------------");
		fmt.Printf("Enter -> \n 1 -> add \n 2 -> subtract \n 3 -> multiply \n 4 -> divide \n 5 -> exit the app\n");
		fmt.Scanf("%d", &input)
		
		if input == 5 {
			break
		}

		fmt.Printf("Enter numbers:\n")

		var n1, n2 int32
		fmt.Scanf("%d", &n1)
		fmt.Println()
		fmt.Scanf("%d", &n2)
		switch input {
			case 1 : fmt.Printf(add(n1, n2))
			case 2 : fmt.Printf(subtract(n1, n2))
			case 3 : fmt.Printf(multiply(n1, n2))
			case 4 : fmt.Printf(divide(n1, n2))
		}
		fmt.Println()
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println()
	}
}

func add(x int32, y int32) (string, int32) {
	var msg string = x + " + " + y + " ="
	res := x+y
	return msg, res
}

func subtract(x int32, y int32) (string, int32) {
	var msg string = x + " - " + y + " ="
	res := x-y
	return msg, res
}

func multiply(x int32, y int32) (string, int32) {
	var msg string = x + " * " + y + " ="
	res := x*y
	return msg, res
}

func divide(x int32, y int32) (string, int32) {
	var msg string = x + " / " + y + " ="
	res := x/y
	return msg, res
}