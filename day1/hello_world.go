package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	name := "rach"
	fmt.Printf("Hello %s, have a great day!\n", name)

	n1, n2 := 10, 20
	fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)

	var inputName string
	fmt.Scanf("%s", &inputName)
	fmt.Printf("Hello %s, have a great day!\n", inputName)
}