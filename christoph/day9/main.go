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
	visited := make(map[string]int)
	readFile, err := os.Open("input.txt")
	checkError(err)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	visited["0_0"] = 1
	xPosHead := 0
	yPosHead := 0
	xPosTail := 0
	yPosTail := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " ")
		dir := split[0]
		steps, err := strconv.Atoi(split[1])
		checkError(err)
		fmt.Printf("=== %s %d\n", dir, steps)
		for i := 0; i < steps; i++ {
			if dir == "U" {
				yPosHead--
			} else if dir == "D" {
				yPosHead++
			} else if dir == "R" {
				xPosHead++
			} else if dir == "L" {
				xPosHead--
			}
			//fmt.Printf("Head moves to %d %d\n", xPosHead, yPosHead)
			if !touching(yPosHead, yPosTail, xPosHead, xPosTail) {
				if xPosHead == xPosTail && yPosTail > yPosHead {
					yPosTail--
				} else if xPosHead == xPosTail && yPosTail < yPosHead {
					yPosTail++
				} else if yPosHead == yPosTail && xPosTail < xPosHead {
					xPosTail++
				} else if yPosHead == yPosTail && xPosTail > xPosHead {
					xPosTail--
				} else if yPosHead > yPosTail && xPosHead > xPosTail {
					xPosTail++
					yPosTail++
				} else if yPosHead < yPosTail && xPosHead < xPosTail {
					xPosTail--
					yPosTail--
				} else if yPosHead > yPosTail && xPosHead < xPosTail {
					xPosTail--
					yPosTail++
				} else if yPosHead < yPosTail && xPosHead > xPosTail {
					xPosTail++
					yPosTail--
				}
				fmt.Printf("Tail moves to %d %d\n", xPosTail, yPosTail)
				pos := fmt.Sprintf("%d_%d", xPosTail, yPosTail)
				visited[pos] = visited[pos] + 1

			}
		}
	}
	fmt.Println(len(visited))
}

func touching(yPosHead int, yPosTail int, xPosHead int, xPosTail int) bool {
	xDiff := xPosHead - xPosTail
	yDiff := yPosHead - yPosTail
	if xDiff < 1 {
		xDiff = xDiff * -1
	}
	if yDiff < 1 {
		yDiff = yDiff * -1
	}
	return xDiff <= 1 && yDiff <= 1
}

func round2() {
	visited := make(map[string]int)
	readFile, err := os.Open("input.txt")
	checkError(err)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	visited["0_0"] = 1
	xPos := make([]int, 10)
	yPos := make([]int, 10)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		split := strings.Split(line, " ")
		dir := split[0]
		steps, err := strconv.Atoi(split[1])
		checkError(err)
		fmt.Printf("=== %s %d\n", dir, steps)
		for i := 0; i < steps; i++ {
			if dir == "U" {
				yPos[0]--
			} else if dir == "D" {
				yPos[0]++
			} else if dir == "R" {
				xPos[0]++
			} else if dir == "L" {
				xPos[0]--
			}
			fmt.Printf("Head moves to %d %d\n", xPos[0], yPos[0])
			for r := 1; r < 10; r++ {
				xPosHead := xPos[r-1]
				yPosHead := yPos[r-1]
				xPosTail := xPos[r]
				yPosTail := yPos[r]
				if !touching(yPosHead, yPosTail, xPosHead, xPosTail) {
					if xPosHead == xPosTail && yPosTail > yPosHead {
						yPos[r]--
					} else if xPosHead == xPosTail && yPosTail < yPosHead {
						yPos[r]++
					} else if yPosHead == yPosTail && xPosTail < xPosHead {
						xPos[r]++
					} else if yPosHead == yPosTail && xPosTail > xPosHead {
						xPos[r]--
					} else if yPosHead > yPosTail && xPosHead > xPosTail {
						xPos[r]++
						yPos[r]++
					} else if yPosHead < yPosTail && xPosHead < xPosTail {
						xPos[r]--
						yPos[r]--
					} else if yPosHead > yPosTail && xPosHead < xPosTail {
						xPos[r]--
						yPos[r]++
					} else if yPosHead < yPosTail && xPosHead > xPosTail {
						xPos[r]++
						yPos[r]--
					}
					fmt.Printf("Part %d moves to %d %d\n", r, xPos[r], yPos[r])
					if r == 9 {
						pos := fmt.Sprintf("%d_%d", xPos[r], yPos[r])
						visited[pos] = visited[pos] + 1
					}
				}
			}
		}
	}
	fmt.Println(len(visited))
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
