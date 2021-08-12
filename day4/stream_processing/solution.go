package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

var sourceChannel chan int = make(chan int)
var evenChannel chan int = make(chan int)
var oddChannel chan int = make(chan int)
var evenSource chan int = make(chan int)
var oddSource chan int = make(chan int)
var wg sync.WaitGroup = sync.WaitGroup{}

func main() {

	go readFileIntoSourceChannel("data1.dat")
	go readFileIntoSourceChannel("data2.dat")
	wg.Add(4)
	go splitter()
	go sum(evenChannel, evenSource)
	go sum(oddChannel, oddSource)
	go merger()
	wg.Wait()
}

func merger() {
	fmt.Println("even sum : ", <-evenSource)
	fmt.Println("odd sum : ", <-oddSource)
	close(evenSource)
	close(oddSource)
	wg.Done()
}

func sum(sourceChannel chan int, destChannel chan int) {
	res := 0
	for num := range sourceChannel {
		res += num
	}
	fmt.Println("SUM IS DONE = ", res)
	destChannel <- res
	close(sourceChannel)
	wg.Done()
}

func splitter() {
	for num := range sourceChannel {
		if num%2 == 0 {
			evenChannel <- num
			fmt.Println("writing into even channel -> ", num)
		} else {
			oddChannel <- num
			fmt.Println("writing into odd channel -> ", num)
		}
	}
	close(sourceChannel)
	wg.Done()
}

func readFileIntoSourceChannel(fileName string) {
	fileReader, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewReader(fileReader)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		num, _ := strconv.Atoi(string(line))
		fmt.Println("line -> ", line, " | writing into the sourceChannel <- ", num)
		sourceChannel <- num
	}
}
