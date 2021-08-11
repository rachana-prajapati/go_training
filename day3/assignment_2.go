package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40))                                         //=> 100
	fmt.Println(sum(10))                                                     //=> 10
	fmt.Println(sum())                                                       //=> 0
	fmt.Println(sum(10, "20", 30, "40"))                                     //=> 100
	fmt.Println(sum(10, "20", 30, "40", "abc"))                              //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                                  //=> 100
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{10, 20}}))           //=> 130
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{10, "20"}})) //=> 130
}

func convertStringToNum(s string) int {
	res := 0
	if num, err := strconv.Atoi(s); err == nil {
		res = num
	}
	return res
}

func sum(nos ...interface{}) int {
	//nos => slice (collection) of int
	res := 0
	for _, val := range nos {

		switch v := val.(type) {
		case int:
			res += v
		case string:
			res += convertStringToNum(v)
		case []int:
			noList := make([]interface{}, len(v))
			for idx, num := range v {
				noList[idx] = num
			}
			res += sum(noList...)
		case []interface{}:
			res += sum(v...)
		}
	}
	return res
}
