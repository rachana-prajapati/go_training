package main

import "fmt"

func main() {
	//arrays
	var arr [5]int = [5]int{5, 6, 7, 3, 6}
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("using the range construct")
	for idx, val := range arr {
		fmt.Println(idx, val)
	}

	fmt.Println("using the range construct with placeholder")
	for _, val := range arr {
		fmt.Println(val)
	}

	//creating the copy of an array
	copyOfArr := arr
	arr[4] = 500
	fmt.Println(copyOfArr, arr)

	//===============================================================================================//

	//slices -> varying size of list of array
	var names []string = []string{"rach", "shaggy", "khloe"}
	/*
		-> If you mention the size, its an array
		-> If you don't mention the size, its a slice
	*/
	fmt.Println(names)

	names = append(names, "sanvi", "bunty")
	fmt.Println(names)

	newNames := []string{"rajesh", "ganesh"}
	names = append(names, newNames...)
	fmt.Println(names)

	//slicing -> get a subset of values
	/*
		[lo : hi] -> lo to hi-1
		[lo :] -> lo to end
		[: hi] -> from the beginning to hi-1
		[:] -> creates copy of slice
	*/
	fmt.Println("[1:3] -> ", names[1:3])
	fmt.Println("[3:] -> ", names[3:])
	fmt.Println("[:3] -> ", names[:3])

	/*
		useful ref:
		- https://research.swtch.com/godata
		- https://blog.golang.org/slices
	*/
	fmt.Println("len(names) -> ", len(names))
	fmt.Println("cap(names) -> ", cap(names))

	names = append(names, "a", "b", "c", "d", "e")
	fmt.Println("len(names) -> ", len(names))
	fmt.Println("cap(names) -> ", cap(names))

	/*
		Q// Can I create a slice with pre-defined capacity as memory allocation is a costly operation?
		-> make
	*/
	list := make([]int, 5, 100) //len -> 5, cap -> 100
	fmt.Println(list, len(list), cap(list))
	list[0] = 1
	list[1] = 10
	list[2] = 20
	list[3] = 30
	list[4] = 40
	fmt.Println(list)
	//list[5] = 50 //gives index out of bound exception/panic
	list = append(list, 50)
	fmt.Println(list)
	fmt.Println(list, len(list), cap(list))

	//===============================================================================================//

	//maps -> key,val collection
	fmt.Println("======================================================================")
	fmt.Println("=================================MAPS==================================")
	fmt.Println("======================================================================")
	var cityRanks map[string]int = map[string]int{
		"blr": 1,
		"bdq": 2,
		"amd": 3,
		"mys": 4,
	}
	fmt.Println(cityRanks)
	fmt.Println("rank of amd -> ", cityRanks["amd"])

	//adding a new entry
	cityRanks["pune"] = 5
	fmt.Println(cityRanks)

	//check if a key exists or not
	if rank, ok := cityRanks["del"]; ok {
		fmt.Println("rank of del is -> ", rank)
	} else {
		fmt.Println("del doesn't exist")
	}

	//remove a key
	delete(cityRanks, "mys")
	fmt.Println("after removing mys -> ", cityRanks)

	//iterate over map
	fmt.Println("Iterating over the map")
	for key, val := range cityRanks {
		fmt.Println(key, val)
	}

	//changing the val
	cityRanks("pune", 10)
	fmt.Println(cityRanks)
}
