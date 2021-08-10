package main

import "fmt"

func main() {
	fn := func(){
		fmt.Printf("fn is invoked..\n")
	}
	exec(fn)

	addProcess := wrapProcess("adding nos", add)
	addProcess(2,3)
	addProcess(20,30)
	addProcess(200,300)
	addProcess(2000,3000)
	
	subtractProcess := wrapProcess("subtracting nos", subtract)
	subtractProcess(300, 200)
	subtractProcess(3000, 2000)
}

func wrapProcess(msg string, fn func(int, int) int) func(x,y int){

	return func(x, y int) {
		fmt.Printf("%s\n", msg)
		res := fn(x,y)
		fmt.Printf("result: %d\n", res)
	}
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