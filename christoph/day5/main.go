package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func main() {
	round2()
}

func round1() {
	stacks, procedures := readInput("input.txt")
	printStacksTopLayer(stacks)
	for _, procedure := range procedures {
		for i := 0; i < procedure.Amount; i++ {
			val := stacks[procedure.From-1].Pop()
			stacks[procedure.To-1].Push(val)
		}
		printStacksTopLayer(stacks)
	}
	printStacksTopLayer(stacks)
}

func printStacksTopLayer(stacks []*stack.Stack) {
	for i := 0; i < len(stacks); i++ {
		val := stacks[i].Peek()
		if val == nil {
			fmt.Printf(" ")
		} else {
			fmt.Printf("%s", val)
		}
	}
	fmt.Printf("\n")

}

func round2() {
	stacks, procedures := readInput("example.txt")
	printStacksTopLayer(stacks)
	for _, procedure := range procedures {
		temp := stack.New()
		for i := 0; i < procedure.Amount; i++ {
			temp.Push(stacks[procedure.From-1].Pop())
		}
		for i := 0; i < procedure.Amount; i++ {
			stacks[procedure.To-1].Push(temp.Pop())
		}
		printStacksTopLayer(stacks)
	}
	printStacksTopLayer(stacks)
}

func read(assignment string) (int, int) {
	eles := strings.Split(assignment, "-")
	start, err := strconv.Atoi(eles[0])
	checkError(err)
	end, err := strconv.Atoi(eles[1])
	checkError(err)
	return start, end
}

type Arrangement struct {
	Amount int
	From   int
	To     int
}

func readInput(file string) ([]*stack.Stack, []Arrangement) {
	stackLines := make([]string, 0)
	stacks := make([]*stack.Stack, 0)
	arrangements := make([]Arrangement, 0)
	readFile, err := os.Open(file)
	checkError(err)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	part2 := false

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			part2 = true
		} else if part2 {
			r := regexp.MustCompile("move ([0-9]*) from ([0-9]*) to ([0-9]*)")
			res := r.FindStringSubmatch(line)
			amount, err := strconv.Atoi(res[1])
			checkError(err)
			from, err := strconv.Atoi(res[2])
			checkError(err)
			to, err := strconv.Atoi(res[3])
			checkError(err)
			arrangements = append(arrangements, Arrangement{
				Amount: amount,
				From:   from,
				To:     to,
			})

		} else {
			if strings.Contains(line, "[") {
				stackLines = append(stackLines, line)
			}
		}
	}

	stacksCreated := false
	for i := len(stackLines) - 1; i > -1; i-- {
		line := stackLines[i]
		if !stacksCreated {
			stacks = make([]*stack.Stack, (len(line)/4)+2)
			stacksCreated = true
			for i := 0; i < len(stacks); i++ {
				stacks[i] = stack.New()
			}
		}
		counter := 0
		for len(line) > 3 {
			if line[:4] != "    " {
				char := getCharacter(line[:4])
				if char != "" {
					stacks[counter].Push(char)
				}
			}
			line = line[4:]
			counter++
		}

		char := getCharacter(line)
		if char != "" {
			stacks[counter].Push(char)
		}
	}
	return stacks, arrangements
}
func getCharacter(input string) string {
	in := strings.TrimSpace(input)
	if in != "" {
		r := regexp.MustCompile(`\[([A-Z])\]`)
		res := r.FindStringSubmatch(in)
		return res[1]
	} else {
		return ""
	}
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
