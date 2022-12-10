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
	signalStrengthSum := 0
	cycles := 0
	x := 1
	lines := readInput("input.txt")
	for _, line := range lines {
		if line == "noop" {
			cycles++
			signalStrengthSum = checkSignalStrength(signalStrengthSum, cycles, x)
		} else if strings.HasPrefix(line, "addx") {
			parts := strings.Split(line, " ")
			v, err := strconv.Atoi(parts[1])
			checkError(err)
			cycles++
			signalStrengthSum = checkSignalStrength(signalStrengthSum, cycles, x)
			cycles++
			signalStrengthSum = checkSignalStrength(signalStrengthSum, cycles, x)
			x = x + v
		} else {
			fmt.Printf("Unrecognized command '%s'\n", line)
		}
	}
	fmt.Printf("Program ran for %d cycles, x is %d, signalStrengthSum %d\n", cycles, x, signalStrengthSum)
}

func round2() {
	display := make([]string, 7)
	display[0] = "                                           "
	display[1] = "                                           "
	display[2] = "                                           "
	display[3] = "                                           "
	display[4] = "                                           "
	display[5] = "                                           "
	display[6] = "                                           "
	cycle := 0
	x := 1

	lines := readInput("input.txt")
	for _, line := range lines {
		if line == "noop" {
			cycle++
			display = draw(display, cycle, x)
		} else if strings.HasPrefix(line, "addx") {
			parts := strings.Split(line, " ")
			v, err := strconv.Atoi(parts[1])
			checkError(err)
			cycle++
			display = draw(display, cycle, x)
			cycle++
			display = draw(display, cycle, x)
			x = x + v
		} else {
			fmt.Printf("Unrecognized command '%s'\n", line)
		}
	}
	for _, line := range display {
		fmt.Println(line)
	}
}

func draw(display []string, cycle int, x int) []string {
	row := 0
	col := cycle
	sprite := x
	if cycle > 0 && cycle <= 40 {
		row = 0
	} else if cycle > 40 && cycle <= 80 {
		row = 1
	} else if cycle > 80 && cycle <= 120 {
		row = 2
	} else if cycle > 120 && cycle <= 160 {
		row = 3
	} else if cycle > 160 && cycle <= 200 {
		row = 4
	} else if cycle > 200 && cycle <= 240 {
		row = 5
	}
	col = cycle - (40 * row) - 1
	//fmt.Printf("============\n")
	//fmt.Printf("Cycle  %d \n", cycle)
	//fmt.Printf("Selected Row %d \n", row)
	//fmt.Printf("Sprite Position %d \n", sprite)
	//fmt.Printf("Selected Column %d \n", col)

	if col == sprite || col == sprite-1 || col == sprite+1 {
		display[row] = display[row][:col] + "#" + display[row][col+1:]
	} else {
		display[row] = display[row][:col] + "." + display[row][col+1:]
	}
	//fmt.Printf("Current CRT row: %s\n", display[row])

	return display
}

func checkSignalStrength(signalStrengthSum int, cycles int, x int) int {
	if cycles == 20 {
		fmt.Printf("SignalStrength for cycles=%d x=%d is %d\n", cycles, x, cycles*x)
		return signalStrengthSum + (cycles * x)
	} else if (cycles-20)%40 == 0 {
		fmt.Printf("SignalStrength for cycles=%d x=%d is %d\n", cycles, x, cycles*x)
		return signalStrengthSum + (cycles * x)
	} else {
		return signalStrengthSum
	}
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
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
