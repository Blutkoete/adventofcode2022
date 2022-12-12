package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	round2()
}

type NodeIndex struct {
	x int
	y int
}

func round1() {
	graph, startPos, endPos, xDim, yDim := readInput("input.txt")
	steps := runDjikstra(graph, startPos, endPos, xDim, yDim)
	fmt.Printf("steps %d\n", steps) // input.txt = 339 example.txt=31
}

func round2() {
	graph, _, endPos, xDim, yDim := readInput("input.txt")
	positionsToCheck := getAllStartPos(graph)
	lowest := 2147483647
	for _, pos := range positionsToCheck {
		steps := runDjikstra(graph, *pos, endPos, xDim, yDim)
		if steps != -1 {
			if steps < lowest {
				lowest = steps
			}
		}
	}
	fmt.Printf("lowest %d\n", lowest) // input.txt = ??? example.txt=29
}

func getAllStartPos(graph [][]rune) []*NodeIndex {
	indizes := make([]*NodeIndex, 0)
	for y, row := range graph {
		for x, char := range row {
			if char == 'a' || char == 'S' {
				indizes = append(indizes, &NodeIndex{
					x: x,
					y: y,
				})
			}
		}
	}
	return indizes
}

func runDjikstra(graph [][]rune, startPos NodeIndex, endPos NodeIndex, xDim int, yDim int) int {
	prev := make([][]*NodeIndex, yDim)
	dist := make([][]int, yDim)
	queue := make([]*NodeIndex, 0)
	for y := 0; y < yDim; y++ {
		dist[y] = make([]int, xDim)
		prev[y] = make([]*NodeIndex, xDim)
		for x := 0; x < xDim; x++ {
			dist[y][x] = 2147483647
			prev[y][x] = nil
			queue = append(queue, &NodeIndex{
				x: x,
				y: y,
			})
		}
	}
	dist[startPos.y][startPos.x] = 0
	for len(queue) > 1 {
		u := nodeWithMinDist(queue, dist)
		if u == nil {
			fmt.Printf("Queue not nil, but nodeWithMinDist not found\n")
			break
		}
		//fmt.Printf("Checking node [%d,%d]\n", u.x, u.y)
		queue = removeFrom(queue, u)
		for _, v := range neighborsStillInQueue(u, queue, graph, xDim, yDim) {
			alt := dist[u.y][u.x] + 1 // check if 1 is correct
			if alt < dist[v.y][v.x] {
				dist[v.y][v.x] = alt
				prev[v.y][v.x] = u
			}
		}
	}

	//fmt.Printf("Ran Djikstra, constructing path\n")
	counter := 1
	finished := false
	currentX := endPos.x
	currentY := endPos.y
	for !finished {
		previous := prev[currentY][currentX]
		if previous == nil {
			return -1
		}
		currentX = previous.x
		currentY = previous.y
		if currentX == startPos.x && currentY == startPos.y {
			finished = true
			break
		}
		counter++
	}
	return counter
}

func neighborsStillInQueue(u *NodeIndex, q []*NodeIndex, graph [][]rune, xDim int, yDim int) []*NodeIndex {
	neighbors := make([]*NodeIndex, 0)
	if u.x-1 >= 0 {
		if elevationCheck(graph, u.x-1, u.y, u.x, u.y) {
			v := NodeIndex{
				x: u.x - 1,
				y: u.y,
			}
			if search(q, &v) != -1 {
				neighbors = append(neighbors, &v)
			}
		} else {
			//fmt.Printf("Node [%d,%d] not reachable from [%d,%d]\n", u.x-1, u.y, u.x, u.y)
		}
	}
	if u.x+1 < xDim {
		if elevationCheck(graph, u.x+1, u.y, u.x, u.y) {
			v := NodeIndex{
				x: u.x + 1,
				y: u.y,
			}
			if search(q, &v) != -1 {
				neighbors = append(neighbors, &v)
			}
		} else {
			//fmt.Printf("Node [%d,%d] not reachable from [%d,%d]\n", u.x+1, u.y, u.x, u.y)
		}
	}
	if u.y-1 >= 0 {
		if elevationCheck(graph, u.x, u.y-1, u.x, u.y) {
			v := NodeIndex{
				x: u.x,
				y: u.y - 1,
			}
			if search(q, &v) != -1 {
				neighbors = append(neighbors, &v)
			}
		} else {
			//fmt.Printf("Node [%d,%d] not reachable from [%d,%d]\n", u.x, u.y-1, u.x, u.y)
		}
	}
	if u.y+1 < yDim {
		if elevationCheck(graph, u.x, u.y+1, u.x, u.y) {
			v := NodeIndex{
				x: u.x,
				y: u.y + 1,
			}
			if search(q, &v) != -1 {
				neighbors = append(neighbors, &v)
			}
		} else {
			//fmt.Printf("Node [%d,%d] not reachable from [%d,%d]\n", u.x, u.y+1, u.x, u.y)
		}
	}
	return neighbors
}

func elevationCheck(graph [][]rune, tX int, tY int, sX int, sY int) bool {
	sourceRuneValue := runeToInt(graph[sY][sX])
	targetRuneValue := runeToInt(graph[tY][tX])
	if targetRuneValue-sourceRuneValue <= 1 {
		return true
	} else {
		return false
	}
}

func runeToInt(r rune) int {
	if r == 'E' {
		return int('z'-'a') + 1
	} else if r == 'S' {
		return int('a'-'a') + 1
	} else {
		return int(r-'a') + 1
	}
}

func search(q []*NodeIndex, u *NodeIndex) int {
	index := -1
	for i := 0; i < len(q); i++ {
		if q[i].x == u.x && q[i].y == u.y {
			index = i
		}
	}
	return index
}

func removeFrom(q []*NodeIndex, u *NodeIndex) []*NodeIndex {
	indexToRemove := search(q, u)
	if indexToRemove != -1 {
		if len(q) == indexToRemove {
			return q[:indexToRemove]
		} else {
			return append(q[:indexToRemove], q[indexToRemove+1:]...)
		}
	}
	return q
}

func nodeWithMinDist(q []*NodeIndex, dist [][]int) *NodeIndex {
	minDist := 2147483648
	targetX := -1
	targetY := -1
	for _, u := range q {
		if dist[u.y][u.x] < minDist {
			minDist = dist[u.y][u.x]
			targetX = u.x
			targetY = u.y
		}
	}
	if targetX == -1 && targetY == -1 {
		return nil
	} else {
		return &NodeIndex{
			x: targetX,
			y: targetY,
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func readInput(file string) ([][]rune, NodeIndex, NodeIndex, int, int) {
	graph := make([][]rune, 0)
	start := NodeIndex{}
	end := NodeIndex{}
	readFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	rowCounter := 0
	colCounter := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		colCounter = len(line)
		row := make([]rune, 0)
		for col, character := range line {
			row = append(row, character)
			if character == 'S' {
				start.x = col
				start.y = rowCounter
			} else if character == 'E' {
				end.x = col
				end.y = rowCounter
			}
		}
		graph = append(graph, row)
		rowCounter++
	}
	return graph, start, end, colCounter, rowCounter
}
