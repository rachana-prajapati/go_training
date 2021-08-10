package main

import "fmt"

func main() {
	inc, dec := getCounter()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
}

func getCounter() (func() int, func() int){
	counter := 0

	inc := func() int{
		counter++
		return counter
	}

	dec := func() int{
		counter--
		return counter
	}

	return inc, dec
}

/*
func main() {

	counter := 0

	getCounter := func() int{
		return counter
	}

	setCounter := func(x int){
		counter = x
	}

	fmt.Println(inc(getCounter, setCounter))
	fmt.Println(inc(getCounter, setCounter))
	fmt.Println(inc(getCounter, setCounter))
	fmt.Println(inc(getCounter, setCounter))

	fmt.Println(dec(getCounter, setCounter))
	fmt.Println(dec(getCounter, setCounter))
	fmt.Println(dec(getCounter, setCounter))
	fmt.Println(dec(getCounter, setCounter))
}

func inc(getCounter func() int, setCounter func(int)) int{
	x := getCounter()
	x += 1
	setCounter(x)
	return x
}

func dec(getCounter func() int, setCounter func(int)) int{
	x := getCounter()
	x -= 1
	setCounter(x)
	return x
}
*/