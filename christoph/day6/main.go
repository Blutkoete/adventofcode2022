package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	round2()
}

func round1() {
	content := readInput("input.txt")
	i := 4
	for ; i < len(content); i++ {
		part := content[i-4 : i]
		if dup_count(part) == 0 {
			break
		}
	}
	fmt.Printf("%d\n", i)
}

func round2() {
	content := readInput("input.txt")
	i := 14
	for ; i < len(content); i++ {
		part := content[i-14 : i]
		if dup_count(part) == 0 {
			break
		}
	}
	fmt.Printf("%d\n", i)
}

func readInput(file string) []byte {
	fileBytes, err := ioutil.ReadFile(file)
	checkError(err)
	return fileBytes
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func dup_count(list []byte) int {
	total := 0
	duplicate_frequency := make(map[byte]int)
	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicate_frequency[item]
		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
			total++
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return total
}
