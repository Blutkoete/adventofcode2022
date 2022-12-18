package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var drawChamberSwitch bool = false

func main() {
	part1("input.txt") // 3105 too high, should be 3048
}

func part1(file string) {
	newRockYOffset := 3
	newRockXOffset := 2
	numberOfRocks := 2022
	rocks := rockGenerator()
	wind := readInput(file)
	chamber := make([][]int, 0)
	currentWindIndex := 0

	for i := 0; i < numberOfRocks; i++ {
		rock := <-rocks

		drawChamber(chamber)
		// add rock
		rockHeight := len(rock)
		rockBottom := getCurrentHeight(chamber) + newRockYOffset
		chamber = padChamberWithEmptySpace(chamber, rockBottom+rockHeight)
		for y, rockLine := range rock {
			for i, ele := range rockLine {
				if ele == 1 {
					chamber[rockBottom+y][i+newRockXOffset] = 1
				}
			}
		}
		drawChamber(chamber)

		collission := false
		for collission == false {
			// apply wind
			var windDirection int
			currentWindIndex, windDirection = getWind(currentWindIndex, wind)
			shiftRock(windDirection, rockBottom, rockHeight, chamber)
			drawChamber(chamber)

			// rock falls down + 1
			collission = fallRock(rockBottom, rockHeight, chamber)
			drawChamber(chamber)
			if collission {
				hardenRock(rockBottom, rockHeight, chamber)
				drawChamber(chamber)
			} else {
				rockBottom = rockBottom - 1
			}
		}
		//fmt.Printf("rockIndex %d currentHeight %d\n", i, getCurrentHeight(chamber))
	}

	//drawChamber(chamber)
	fmt.Printf("currentHeight %d\n", getCurrentHeight(chamber))
}

func padChamberWithEmptySpace(chamber [][]int, targetHeight int) [][]int {
	for y := len(chamber); y < targetHeight; y++ {
		chamber = append(chamber, make([]int, 7))
	}
	return chamber
}

func getCurrentHeight(chamber [][]int) int {
	for i, row := range chamber {
		rowIsEmpty := true
		for x := 0; x < len(row); x++ {
			if row[x] == 2 {
				rowIsEmpty = false
			}
		}
		if rowIsEmpty {
			return i
		}
	}
	return 0
}

func hardenRock(rockBottom int, rockHeight int, chamber [][]int) {
	for y := rockBottom; y < rockBottom+rockHeight; y++ {
		for x := 0; x < 7; x++ {
			if chamber[y][x] == 1 {
				chamber[y][x] = 2
			}
		}
	}
}

func fallRock(rockBottom int, rockHeight int, chamber [][]int) bool {
	if drawChamberSwitch {
		fmt.Printf("Rock falls\n")
	}
	for y := rockBottom; y < rockBottom+rockHeight; y++ {
		for x := 0; x < 7; x++ {
			if chamber[y][x] == 1 {
				if y-1 == -1 { // collission with ground
					return true
				} else if chamber[y-1][x] == 2 { // collission with rock
					return true
				}
			}
		}
	}

	for y := rockBottom; y < rockBottom+rockHeight; y++ {
		for x := 0; x < 7; x++ {
			if chamber[y][x] == 1 {
				chamber[y-1][x] = 1
				chamber[y][x] = 0
			}
		}
	}
	return false
}

func shiftRock(windDirection int, rockBottom int, rockHeight int, chamber [][]int) {
	// windDirection > = 1, < = 2
	if drawChamberSwitch {
		if windDirection == 1 {
			fmt.Println("wind to right")
		} else {
			fmt.Println("wind to left")
		}
	}
	// Check Collission
	for y := 0; y < rockHeight; y++ {
		if windDirection == 1 {
			if chamber[rockBottom+y][6] == 1 {
				if drawChamberSwitch {
					fmt.Println("Rock blocked by wall")
				}
				return
			}
			for x := 0; x < 6; x++ {
				if chamber[rockBottom+y][x] == 1 && chamber[rockBottom+y][x+1] == 2 {
					if drawChamberSwitch {
						fmt.Println("Rock blocked by rock")
					}
					return
				}
			}
		} else {
			if chamber[rockBottom+y][0] == 1 {
				if drawChamberSwitch {
					fmt.Println("Rock blocked by wall")
				}
				return
			}
			for x := 6; x > 0; x-- {
				if chamber[rockBottom+y][x] == 1 && chamber[rockBottom+y][x-1] == 2 {
					if drawChamberSwitch {
						fmt.Println("Rock blocked by rock")
					}
					return
				}
			}
		}
	}

	// No Collission
	for y := 0; y < rockHeight; y++ {
		if windDirection == 1 {
			for x := 6; x >= 0; x-- {
				if chamber[rockBottom+y][x] == 1 {
					chamber[rockBottom+y][x] = 0
					chamber[rockBottom+y][x+1] = 1
				}
			}
		} else {
			for x := 1; x < 7; x++ {
				if chamber[rockBottom+y][x] == 1 {
					chamber[rockBottom+y][x] = 0
					chamber[rockBottom+y][x-1] = 1
				}
			}
		}
	}
	return
}

func getWind(index int, windDirections []int) (int, int) {
	i := index + 1
	if i == len(windDirections) {
		i = 0
	}
	return i, windDirections[index]
}

func drawChamber(chamber [][]int) {
	if drawChamberSwitch {
		drawHeight := len(chamber) - 20
		if drawHeight < 0 {
			drawHeight = 0
		}
		fmt.Println("      =========")
		for i := len(chamber) - 1; i >= drawHeight; i-- {
			fmt.Printf("%6d|", i)
			for _, v := range chamber[i] {
				if v == 0 {
					fmt.Print(".")
				} else if v == 1 {
					fmt.Print("@")
				} else if v == 2 {
					fmt.Print("#")
				}
			}
			fmt.Printf("|\n")
		}
		fmt.Print("      +-------+")
		fmt.Println("")
	}
}

func rockGenerator() chan [][]int {
	rocks := make([]string, 0)
	readFile, err := os.Open("rocks.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	currentRock := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			rocks = append(rocks, currentRock)
			currentRock = ""
		}
		currentRock = fmt.Sprintf("%s\n%s", currentRock, line)
	}
	rocks = append(rocks, currentRock)

	channel := make(chan [][]int)
	go func() {
		for {
			for _, rock := range rocks {
				lines := make([]string, 0)
				fileScanner := bufio.NewScanner(strings.NewReader(rock))
				fileScanner.Split(bufio.ScanLines)
				for fileScanner.Scan() {
					line := fileScanner.Text()
					if line != "" {
						lines = append(lines, line)
					}
				}
				currentRock := make([][]int, 0)
				for i := len(lines) - 1; i >= 0; i-- {
					currentRockLine := make([]int, 0)
					currentLine := lines[i]
					for x := 0; x < len(currentLine); x++ {
						if currentLine[x] == '#' {
							currentRockLine = append(currentRockLine, 1)
						} else {
							currentRockLine = append(currentRockLine, 0)
						}
					}
					currentRock = append(currentRock, currentRockLine)
				}
				channel <- currentRock
			}
		}
	}()
	return channel
}

func readInput(file string) []int {
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

	wind := make([]int, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		if lines[0][i] == '>' {
			wind[i] = 1
		} else if lines[0][i] == '<' {
			wind[i] = 2
		}
	}
	return wind
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
