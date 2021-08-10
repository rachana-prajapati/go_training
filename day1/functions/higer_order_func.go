package main

import "fmt"

func main() {
	add := func(x,y int) int {
		return x+y
	}
	fmt.Println(add(100, 200))

	//anonymous func
	func(){
		fmt.Printf("anonymous func invoked..\n")
	}()

	func(x,y int) {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}(100, 200)
}