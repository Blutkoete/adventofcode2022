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
	/*timeUnits, sandUnits := round1("example.txt", 500, 0)
	fmt.Printf("1 example.txt Timeunits %d Sandunits %d\n", timeUnits, sandUnits)
	timeUnits, sandUnits = round1("input.txt", 500, 0)
	fmt.Printf("1 input.txt Timeunits %d Sandunits %d\n", timeUnits, sandUnits)*/
	timeUnits, sandUnits := round2("example.txt", 500, 0)
	fmt.Printf("2 example.txt Timeunits %d Sandunits %d\n", timeUnits, sandUnits)
	timeUnits, sandUnits = round2("input.txt", 500, 0)
	fmt.Printf("2 input.txt Timeunits %d Sandunits %d\n", timeUnits, sandUnits)
}

func round1(file string, sandSourceX int, sandSourceY int) (int, int) {
	lines := readInput(file)
	caveMap := toMap(lines, 600, 600, false)
	timeUnits := 0
	sandUnits := 0
	done := false
	for !done {
		currentSandPositionX := sandSourceX
		currentSandPositionY := sandSourceY + 1
		for {
			//fmt.Println("=======================================")
			//drawMapWithFallingGrain(caveMap, sandSourceX, sandSourceY, currentSandPositionX, currentSandPositionY, 493, 504, 0, 10)
			//fmt.Println("=======================================")
			collision := checkCollision(caveMap, currentSandPositionX, currentSandPositionY)
			if collision == 0 {
				currentSandPositionY++
			} else if collision == 1 {
				currentSandPositionY++
				currentSandPositionX--
			} else if collision == 2 {
				currentSandPositionY++
				currentSandPositionX++
			} else if collision == 3 {
				caveMap[currentSandPositionY][currentSandPositionX] = 2
				sandUnits++
				break
			} else if collision == -1 {
				done = true
				break
			}
			timeUnits++
		}
	}
	return timeUnits, sandUnits
}

func round2(file string, sandSourceX int, sandSourceY int) (int, int) {
	lines := readInput(file)
	caveMap := toMap(lines, 1200, 600, true)
	timeUnits := 0
	sandUnits := 0
	done := false
	for !done {
		currentSandPositionX := sandSourceX
		currentSandPositionY := sandSourceY
		for {
			//fmt.Println("=======================================")
			//drawMapWithFallingGrain(caveMap, sandSourceX, sandSourceY, currentSandPositionX, currentSandPositionY, 493, 506, 0, 14)
			//fmt.Println("=======================================")
			collision := checkCollision(caveMap, currentSandPositionX, currentSandPositionY)
			if collision == 0 {
				currentSandPositionY++
			} else if collision == 1 {
				currentSandPositionY++
				currentSandPositionX--
			} else if collision == 2 {
				currentSandPositionY++
				currentSandPositionX++
			} else if collision == 3 {
				caveMap[currentSandPositionY][currentSandPositionX] = 2
				sandUnits++
				if currentSandPositionY == 0 && currentSandPositionX == 500 {
					done = true
					break
				}
				break
			} else if collision == -1 {
				done = true
				break
			}
			timeUnits++
		}
	}
	return timeUnits, sandUnits
}

func checkCollision(caveMap [][]int, currentSandPositionX int, currentSandPositionY int) int {
	if len(caveMap) == currentSandPositionY+1 {
		return -1
	}
	if caveMap[currentSandPositionY+1][currentSandPositionX] == 0 {
		return 0 // no collision
	} else if caveMap[currentSandPositionY+1][currentSandPositionX] == 1 || caveMap[currentSandPositionY+1][currentSandPositionX] == 2 {
		// collision
		if caveMap[currentSandPositionY+1][currentSandPositionX-1] == 1 || caveMap[currentSandPositionY+1][currentSandPositionX-1] == 2 {
			// left collision, check right side
			if caveMap[currentSandPositionY+1][currentSandPositionX+1] == 1 || caveMap[currentSandPositionY+1][currentSandPositionX+1] == 2 {
				// right collision
				return 3
			} else {
				return 2
			}
		} else {
			return 1
		}
	}
	return -1
}

func toMap(lines []string, xDim int, yDim int, hasFloor bool) [][]int {
	caveMap := make([][]int, yDim)
	for i := 0; i < len(caveMap); i++ {
		caveMap[i] = make([]int, xDim)
	}

	heighestRow := -1
	for _, line := range lines {
		elements := strings.Split(line, " -> ")
		var currentRow int = -1
		var currentCol int = -1
		for _, element := range elements {
			parts := strings.Split(element, ",")
			colString := parts[0]
			rowString := parts[1]
			if currentRow == -1 {
				currentCol, _ = strconv.Atoi(colString)
				currentRow, _ = strconv.Atoi(rowString)
				if currentRow > heighestRow {
					heighestRow = currentRow
				}
			} else {
				targetCol, _ := strconv.Atoi(colString)
				targetRow, _ := strconv.Atoi(rowString)

				var fromX, toX int
				if currentCol > targetCol {
					fromX = targetCol
					toX = currentCol
				} else {
					fromX = currentCol
					toX = targetCol
				}
				var fromY, toY int
				if currentRow > targetRow {
					fromY = targetRow
					toY = currentRow
				} else {
					fromY = currentRow
					toY = targetRow
				}
				for y := fromY; y <= toY; y++ {
					for x := fromX; x <= toX; x++ {
						caveMap[y][x] = 1
					}
				}
				currentRow = targetRow
				currentCol = targetCol
				if currentRow > heighestRow {
					heighestRow = currentRow
				}
			}
		}
	}

	if hasFloor {
		for x := 0; x < len(caveMap[heighestRow+2]); x++ {
			caveMap[heighestRow+2][x] = 1
		}
	}
	return caveMap
}

func drawMap(caveMap [][]int, sandSourceX int, sandSourceY int, startX int, endX int, startY int, endY int) {
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			if x == sandSourceX && y == sandSourceY {
				fmt.Printf("+") // sandsource
			} else if caveMap[y][x] == 0 {
				fmt.Printf(".") // air
			} else if caveMap[y][x] == 1 {
				fmt.Printf("#") // rock
			} else if caveMap[y][x] == 2 {
				fmt.Printf("o") // sand_at_rest
			}
		}
		fmt.Println("")
	}
}

func drawMapWithFallingGrain(caveMap [][]int, sandSourceX int, sandSourceY int, grainX int, grainY int, startX int, endX int, startY int, endY int) {
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			if x == sandSourceX && y == sandSourceY {
				fmt.Printf("+") // sandsource
			} else if x == grainX && y == grainY {
				fmt.Printf("s") // sandsource
			} else if caveMap[y][x] == 0 {
				fmt.Printf(".") // air
			} else if caveMap[y][x] == 1 {
				fmt.Printf("#") // rock
			} else if caveMap[y][x] == 2 {
				fmt.Printf("o") // sand_at_rest
			}
		}
		fmt.Println("")
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
