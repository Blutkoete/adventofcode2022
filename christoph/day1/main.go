package main

import (
	"bufio"
	"fmt"
	"sort"
	"os"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")
 
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	index := 1
	current := 0
	var elfs []int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			fmt.Printf("elf %d has %d calories\n",index,current)
			elfs = append(elfs, current)
			current = 0
			index = index +1
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			current = current + i
		}
	}
	sort.Ints(elfs)
	fmt.Printf("%v %d\n",elfs[len(elfs)-3:], sum(elfs[len(elfs)-3:]))
	readFile.Close()
}

func sum(array []int) int {  
	result := 0  
	for _, v := range array {  
	 result += v  
	}  
	return result  
   }  