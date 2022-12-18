package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("example.txt %d\n", round1("example.txt"))
}

func round1(file string) int {
	lines := readInput(file)
	adjacency, pressure := createGraph(lines)
	releasedValves := make([]string, 0)
	currentValve := "AA"
	currentReleasedPressure := 0
	totalReleasePressure := 0
	currentTime := 1
	maxTime := 30

	for currentTime = 1; currentTime < maxTime; currentTime++ {
		fmt.Printf("== Minute %d ==\n", currentTime)
		remainingTime := maxTime - currentTime
		if currentReleasedPressure > 0 {
			fmt.Printf("Valve %s is open, releasing %d pressure.\n", strings.Join(releasedValves, ","), currentReleasedPressure)
			totalReleasePressure = totalReleasePressure + currentReleasedPressure
		}

		nextValve := nextValve(remainingTime, currentValve, adjacency, pressure, releasedValves)
		if nextValve == "" {
			releasedValves = append(releasedValves, currentValve)
			currentReleasedPressure = currentReleasedPressure + pressure[currentValve]
			fmt.Printf("Opening valve %s, now releasing %d pressure \n", currentValve, currentReleasedPressure)
		} else {
			currentValve = nextValve
			fmt.Printf("Move to %s\n", nextValve)
		}
	}

	return totalReleasePressure

}

func nextValve(remainingTime int, currentValve string, adjacency map[string][]string, pressure map[string]int, releasedValves []string) string {
	maxPressure := -1
	next := ""
	for valve, _ := range adjacency {
		if !contains(releasedValves, valve) {
			timeToReach, n := djikstraDistance(adjacency, currentValve, valve)
			potentialRelease := (remainingTime - timeToReach) * pressure[valve]
			if potentialRelease > maxPressure {
				maxPressure = (remainingTime - timeToReach) * pressure[valve]
				next = n
			}
		}
	}
	return next
}

func contains(array []string, key string) bool {
	for _, key2 := range array {
		if key2 == key {
			return true
		}
	}
	return false
}

func createGraph(lines []string) (map[string][]string, map[string]int) {
	adjacency := make(map[string][]string)
	pressure := make(map[string]int)

	for _, line := range lines {
		parts := strings.Split(line, ";")

		re := regexp.MustCompile("Valve ([A-Z]{2}) has flow rate=([0-9]*)")
		matches := re.FindStringSubmatch(parts[0])
		valve := matches[1]
		valvePressureRelease, err := strconv.Atoi(matches[2])
		checkError(err)

		var targets []string
		if strings.HasPrefix(parts[1], " tunnel leads to valve ") {
			targets = strings.Split(parts[1][23:], ",")
		} else {
			targets = strings.Split(parts[1][24:], ",")
		}
		for i := 0; i < len(targets); i++ {
			targets[i] = strings.TrimSpace(targets[i])
		}
		adjacency[valve] = targets
		pressure[valve] = valvePressureRelease
	}

	return adjacency, pressure
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

func djikstraDistance(adjacency map[string][]string, current string, target string) (int, string) {
	if current == target {
		return 0, ""
	}
	dist := make(map[string]int, 0)
	prev := make(map[string]string, 0)
	queue := make([]string, 0)
	for k, _ := range adjacency {
		dist[k] = 2147483647
		prev[k] = ""
		queue = append(queue, k)
	}
	dist[current] = 0
	sort.Strings(queue)
	for len(queue) > 1 {
		u := nodeWithMinDist(queue, dist)
		if u == "" {
			fmt.Printf("Queue not nil, but nodeWithMinDist not found\n")
			break
		}
		queue = removeFrom(queue, u)
		for _, v := range neighborsStillInQueue(u, queue, adjacency) {
			alt := dist[u] + 1 // check if 1 is correct
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
			}
		}
	}
	cur := target
	for {
		if prev[cur] == current {
			return dist[target], cur
		} else {
			cur = prev[cur]
		}
	}
}

func nodeWithMinDist(q []string, d map[string]int) string {
	currentU := ""
	minDist := 2147483647
	sort.Strings(q)
	for _, u := range q {
		if d[u] < minDist {
			currentU = u
			minDist = d[u]
		}
	}
	return currentU
}

func search(q []string, u string) int {
	index := -1
	for i := 0; i < len(q); i++ {
		if q[i] == u {
			index = i
		}
	}
	return index
}

func removeFrom(q []string, u string) []string {
	indexToRemove := search(q, u)
	if indexToRemove != -1 {
		if len(q) == indexToRemove {
			return q[:indexToRemove]
		} else {
			return append(q[:indexToRemove], q[indexToRemove+1:]...)
		}
	}
	sort.Strings(q)
	return q
}

func neighborsStillInQueue(u string, q []string, adjecency map[string][]string) []string {
	neighbors := make([]string, 0)
	adjecenctNeighbors := adjecency[u]
	sort.Strings(adjecenctNeighbors)
	for _, v := range adjecenctNeighbors {
		if contains(q, v) {
			neighbors = append(neighbors, v)
		}
	}
	sort.Strings(neighbors)
	return neighbors
}
