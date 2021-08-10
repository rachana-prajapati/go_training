package main

import (
	"demo_module/calc"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	x, y := 10, 20
	fmt.Println("AddApi(x,y) -> ", calc.AddApi(x, y))
	fmt.Println("AddApi(x,y) -> ", calc.AddApi(x, y))
	fmt.Println("GetCounter() -> ", calc.GetCounter())

	//adding  a third party dependency
	/*
		run the following command in the module/app folder
		>> go get github.com/fatih/color
			--> downloads the binary and place it under GOPATH/pkg
			--> Also, it adds a dependency in the go.mod file of the module
	*/
	color.Red("This is a red line.")
	color.Blue("This is a blue line.")

	/*
		Build the module
		==============
			>> go build .
				--> creates a binary of the name <module_name>
			>> go build -o <binary_name>
				--> creates a binary of the name <binary_name>


		Run the module using the binary created
		=============================================
			>> ./<binary_name>
				--> executes the main method of the main package
	*/

	/*
		>> go mod vendor
			--> downloads the third party dependeecies in the app/module
	*/
}
