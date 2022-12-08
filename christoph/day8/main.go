package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	round2()
}

func round1() {
	trees := readInput("input.txt")
	_ = trees
	totalVisible := 0
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees); x++ {
			if isVisible(trees, x, y) {
				totalVisible++
			}
		}
	}
	fmt.Printf("%d\n", totalVisible)
}
func round2() {
	trees := readInput("input.txt")
	_ = trees
	heighestScore := 0
	for y := 0; y < len(trees); y++ {
		for x := 0; x < len(trees); x++ {
			score := scenicScore(trees, x, y)
			if heighestScore < score {
				heighestScore = score
			}
		}
	}
	fmt.Printf("%d\n", heighestScore)
}

func isVisible(trees [][]int, x int, y int) bool {
	dimY := len(trees)
	dimX := len(trees[0])
	height := trees[y][x]
	if x == 0 || y == 0 || y == dimY-1 || x == dimX-1 {
		return true
	} else {
		leftVisible := true
		rightVisible := true
		topVisible := true
		bottomVisible := true
		for ix := x - 1; ix >= 0; ix-- {
			if trees[y][ix] >= height {
				leftVisible = false
			}
		}
		for ix := x + 1; ix < dimX; ix++ {
			if trees[y][ix] >= height {
				rightVisible = false
			}
		}
		for iy := y - 1; iy >= 0; iy-- {
			if trees[iy][x] >= height {
				topVisible = false
			}
		}
		for iy := y + 1; iy < dimY; iy++ {
			if trees[iy][x] >= height {
				bottomVisible = false
			}
		}
		return leftVisible || rightVisible || topVisible || bottomVisible
	}
}

func scenicScore(trees [][]int, x int, y int) int {
	dimY := len(trees)
	dimX := len(trees[0])
	height := trees[y][x]
	leftScore := 0
	rightScore := 0
	topScore := 0
	bottomScore := 0
	for ix := x - 1; ix >= 0; ix-- {
		if trees[y][ix] >= height {
			leftScore++
			break
		} else {
			leftScore++
		}
	}
	for ix := x + 1; ix < dimX; ix++ {
		if trees[y][ix] >= height {
			rightScore++
			break
		} else {
			rightScore++
		}
	}
	for iy := y - 1; iy >= 0; iy-- {
		if trees[iy][x] >= height {
			topScore++
			break
		} else {
			topScore++
		}
	}
	for iy := y + 1; iy < dimY; iy++ {
		if trees[iy][x] >= height {
			bottomScore++
			break
		} else {
			bottomScore++
		}
	}
	return leftScore * rightScore * bottomScore * topScore
}
func readInput(file string) [][]int {
	trees := make([][]int, 0)
	readFile, err := os.Open(file)
	checkError(err)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	current := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		row := make([]int, len(line))
		for i, _ := range line {
			tree := line[i : i+1]
			val, err := strconv.Atoi(tree)
			checkError(err)
			row[i] = val
		}
		trees = append(trees, row)
		current++
	}
	return trees
}
func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
