package main

import "fmt"

func main(){
	f1()
}

func f1(){
	defer func(){
		fmt.Printf("f1 execution completed.\n")
	}()
	defer fmt.Printf("another deferred func call.\n")
	fmt.Printf("f1 is invoked.\n")
}