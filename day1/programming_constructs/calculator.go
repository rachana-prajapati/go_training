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
			case 1 : fmt.Printf("%d + %d = %d", n1, n2, n1+n2)
			case 2 : fmt.Printf("%d - %d = %d", n1, n2, n1-n2)
			case 3 : fmt.Printf("%d * %d = %d", n1, n2, n1*n2)
			case 4 : fmt.Printf("%d / %d = %d", n1, n2, n1/n2)
		}
		fmt.Println()
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println()
	}
}