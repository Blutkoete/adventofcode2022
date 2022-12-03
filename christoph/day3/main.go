package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	round2()
}

func round1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	priorities := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		compartment1 := line[:len(line)/2]
		compartment2 := line[len(line)/2:]
		var c1 byte
		var c2 byte
		for i := 0; i < len(compartment1); i++ {
			c1 = compartment1[i]
			for j := 0; j < len(compartment2); j++ {
				c2 = compartment2[j]
				if c1 == c2 {
					break
				}
			}
			if c1 == c2 {
				break
			}
		}
		r := rune(c1)
		r_val := int(r)
		if r_val >= 97 && r_val <= 122 {
			r_val = r_val - 96
		} else if r_val >= 65 && r_val <= 90 {
			r_val = r_val - 64 + 26
		}
		priorities = priorities + r_val
		fmt.Printf("%d %v %s\n", r_val, rune(c1), string(c1))
	}
	fmt.Printf("Score %d\n", priorities)
}

func round2() {
	priorities := 0
	lines := make([]string, 0)
	readFile, err := os.Open("input.txt")
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

	i := 0
	for i < len(lines) {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			c := line[j]
			fmt.Printf("1 %s\n2 %s \n3 %s\n", line, lines[i+1], lines[i+2])
			if strings.ContainsRune(lines[i+1], rune(c)) && strings.ContainsRune(lines[i+2], rune(c)) {
				r := rune(c)
				r_val := int(r)
				if r_val >= 97 && r_val <= 122 {
					r_val = r_val - 96
				} else if r_val >= 65 && r_val <= 90 {
					r_val = r_val - 64 + 26
				}
				priorities = priorities + r_val
				break
			}
		}
		i = i + 3
	}
	fmt.Printf("Score %d\n", priorities)
}
