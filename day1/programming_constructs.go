package main

import "fmt"

func main() {
	//if, for, switch

	num := 20
	if num%2 == 0 {
		fmt.Println(num, " -> is an even number")
	} else {
		fmt.Println(num, " -> is an odd number")
	}

	//******************************************************************************************************//

	/* for */
	sum := 0
	for i:=0; i<10; i++ {
		sum +=i
	}
	fmt.Println("sum : ", sum)

	//imitating the while construct
	whilesum := 1
	for whilesum<45 {
		whilesum += whilesum
	}
	fmt.Println("whilesum : ", whilesum)

	//indefinite loop
	x := 1
	for {
		x += x
		if x > 100 {
			break
		}
	}
	fmt.Println("x : ", x)

	//******************************************************************************************************//

	/* switch */
	newNo := 22
	
	//option#1
	switch newNo%2 {
		case 0:
			fmt.Println(newNo, " is an EVEN number")
		case 1:
			fmt.Println(newNo, " is an ODD number")
	}

	//option#2
	switch  {
		case newNo%2 == 0:
			fmt.Println(newNo, " is an EVEN number")
		case newNo%2 == 1:
			fmt.Println(newNo, " is an ODD number")
	}

	//switch with default
	/*
		Rank by score
		--------------
		score 0-3 ==> "Terrible"
		score 4-7 ==> "Not Bad"
		score 8-9 ==> "Good"
		score 10 ==> "Excellent"
		otherwise ==> "Unknown Score"
	*/
	score := 10
	switch score{
		case 0,1,2,3 : fmt.Println("Terrible") 
		case 4,5,6,7 : fmt.Println("Not Bad") 
		case 8,9 : fmt.Println("Good") 
		case 10 : fmt.Println("Excellent") 
		default : fmt.Println("Unknown Score") 
	}

	//fallthrough construct
	n := 4
	switch n {
		case 0:
			fmt.Println("is zero")
		case 1:
			fmt.Println("n <= 1")
			fallthrough
		case 2:
			fmt.Println("n <= 2")
			fallthrough
		case 3:
			fmt.Println("n <= 3")
			fallthrough
		case 4:
			fmt.Println("n <= 4")
			fallthrough
		case 5:
			fmt.Println("n <= 5")
			fallthrough
		case 6:
			fmt.Println("n <= 6")
			fallthrough
		case 7:
			fmt.Println("n <= 7")
			fallthrough
		case 8:
			fmt.Println("n <= 8")
	}
}