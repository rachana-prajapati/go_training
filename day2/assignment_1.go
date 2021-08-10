package main

import (
	"fmt"
	"strings"
)

/*
Given the following string
Pariatur nisi nostrud id ipsum esse minim velit Lorem eiusmod est reprehenderit sint Esse non proident labore dolore labore nisi dolore dolor proident voluptate minim laborum Ad irure ut Lorem exercitation Ipsum reprehenderit consequat veniam non reprehenderit ut eiusmod magna magna aliquip ut sint Mollit irure fugiat nostrud non excepteur aliquip cillum Labore dolore deserunt irure ea temporConsequat ea quis ipsum esse minim reprehenderit amet Qui do dolor do anim occaecat culpa commodo sunt mollit excepteur laboris tempor Ullamco labore quis culpa laborum sunt Culpa incididunt ad consectetur ut deserunt tempor proident ut Qui quis aliqua fugiat sunt dolore commodo reprehenderit veniam tempor Voluptate sit laboris sunt et do sint ea irure veniam duis laborum Quis irure id officia inNisi velit proident cupidatat laborum velit enim qui deserunt consectetur Ad ea quis veniam cupidatat tempor sunt elit velit Sit anim irure sunt ut exercitation excepteur elit laboris occaecat dolor Duis ad commodo ut quis nisi pariatur anim Fugiat ad id anim velit labore

Find out the size of the word that occurs the most (by size)

for ex:
5 letters occurs 10 times
4 letters occurs 9 times
3 letters occurs 12 times

output : 3 letter words with 12 occurances
*/
func main() {
	msg := "Pariatur nisi nostrud id ipsum esse minim velit Lorem eiusmod est reprehenderit sint Esse non proident labore dolore labore nisi dolore dolor proident voluptate minim laborum Ad irure ut Lorem exercitation Ipsum reprehenderit consequat veniam non reprehenderit ut eiusmod magna magna aliquip ut sint Mollit irure fugiat nostrud non excepteur aliquip cillum Labore dolore deserunt irure ea temporConsequat ea quis ipsum esse minim reprehenderit amet Qui do dolor do anim occaecat culpa commodo sunt mollit excepteur laboris tempor Ullamco labore quis culpa laborum sunt Culpa incididunt ad consectetur ut deserunt tempor proident ut Qui quis aliqua fugiat sunt dolore commodo reprehenderit veniam tempor Voluptate sit laboris sunt et do sint ea irure veniam duis laborum Quis irure id officia inNisi velit proident cupidatat laborum velit enim qui deserunt consectetur Ad ea quis veniam cupidatat tempor sunt elit velit Sit anim irure sunt ut exercitation excepteur elit laboris occaecat dolor Duis ad commodo ut quis nisi pariatur anim Fugiat ad id anim velit labore"
	words := strings.Split(msg, " ")

	//iterate over words and create a map of <len(word), # of occurrances>
	var wordMap map[int]int = getWordLengthToOccurrancesMap(words)

	//get maximum occurring word
	maxOccurrance, lengthOfWord := getMaxOccurringWordByLength(wordMap)
	fmt.Printf("maximum occurring word is having length -> %d, # of occurrances -> %d\n", lengthOfWord, maxOccurrance)
}

func getWordLengthToOccurrancesMap(words []string) map[int]int {
	var wordMap map[int]int = map[int]int{}
	for _, word := range words {
		wordMap[len(word)]++
	}
	return wordMap
}

func getMaxOccurringWordByLength(wordMap map[int]int) (maxOccurrance int, lengthOfWord int) {
	maxOccurrance = 0
	for lengthOfCurrentWord, occurrance := range wordMap {
		if lengthOfCurrentWord > maxOccurrance {
			maxOccurrance = occurrance
			lengthOfWord = lengthOfCurrentWord
		}
	}
	return maxOccurrance, lengthOfWord
}
