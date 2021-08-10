package main

import "fmt"

func main() {
	fmt.Println(sum(10, 20, 30))
}

func sum (nums... int) int {

	//nums => slice (collection) of int
	res := 0
	for i:=0; i<len(nums); i++ {
		res += nums[i]
	}
	return res
}