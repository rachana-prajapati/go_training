package main

import "fmt"

func main() {
	fn := func(){
		fmt.Printf("fn is invoked..\n")
	}
	exec(fn)

	process(add, 2,3)
	process(subtract, 2, 3)
}

func process(fn func(int, int) int, x,y int){
	fmt.Printf("start processing..\n")
	res := fn(x,y)
	fmt.Printf("result: %d\n", res)
	fmt.Printf("done processing..\n")
}

func add(x, y int) int{
	return x+y
}

func subtract(x, y int) int{
	return x-y
}

func exec(fn func()) {
	fmt.Printf("executing function..\n")
	fn()
	fmt.Printf("finished executing func..\n")
}