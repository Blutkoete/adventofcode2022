package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	round2()
}

func round1() {
	contained := 0
	lines := readInput("input.txt")
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		e1start, e1end := read(assignments[0])
		e2start, e2end := read(assignments[1])
		if e1start >= e2start && e1end <= e2end {
			contained = contained + 1
		} else if e2start >= e1start && e2end <= e1end {
			contained = contained + 1
		}
	}
	fmt.Printf("Result %d\n", contained)
}

func round2() {
	contained := 0
	lines := readInput("input.txt")
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		e1start, e1end := read(assignments[0])
		e2start, e2end := read(assignments[1])
		fmt.Printf("%s", line)
		if e1start >= e2start && e1start <= e2end {
			contained = contained + 1
			fmt.Printf(" Overlap start e1\n")
		} else if e1end >= e2start && e1end <= e2end {
			contained = contained + 1
			fmt.Printf(" Overlap end e1\n")
		} else if e2start >= e1start && e2start <= e1end {
			contained = contained + 1
			fmt.Printf(" Overlap start e2\n")
		} else if e2end >= e1start && e2end <= e1end {
			contained = contained + 1
			fmt.Printf(" Overlap end e2\n")
		} else {
			fmt.Printf(" No Overlap\n")
		}
	}
	fmt.Printf("Result %d\n", contained)
}

func read(assignment string) (int, int) {
	eles := strings.Split(assignment, "-")
	start, err := strconv.Atoi(eles[0])
	checkError(err)
	end, err := strconv.Atoi(eles[1])
	checkError(err)
	return start, end
}

func readInput(file string) []string {
	lines := make([]string, 0)
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
