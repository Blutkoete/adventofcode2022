package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func main() {
	round2("input.txt") // not 1752 (too high), 1725 also wrong
}

func round2(file string) {
	lines := readInput(file)
	adjacency, pressure := createGraph(lines)
	reducedAdjecencyList := reduceGraph("AA", adjacency, pressure)

	valves := relevantValves(reducedAdjecencyList, pressure)
	maxPressure := 0
	mu := sync.Mutex{}
	permutations := CombinationGenerator(valves, reducedAdjecencyList)

	for i := 0; i < 50; i++ {
		for permutation := range permutations {
			part1 := permutation[:len(permutation)/2]
			part2 := permutation[len(permutation)/2:]
			p1 := getPressure("AA", part1, pressure, reducedAdjecencyList, 26)
			p2 := getPressure("AA", part2, pressure, reducedAdjecencyList, 26)
			p := p1 + p2
			mu.Lock()
			if p > maxPressure {
				fmt.Printf("New Max Pressure %s %d\n", strings.Join(permutation, ","), p)
				maxPressure = p
			}
			mu.Unlock()
		}
	}
	fmt.Printf("Pressure %d\n", maxPressure)
}

func getPressure(start string, valves []string, pressure map[string]int, graph map[string][]Edge, steps int) int {
	prev := start
	time := steps
	p := 0
	i := 0
	for time > 0 && i < len(valves) {
		nextValve := valves[i]
		//fmt.Printf("Opening Valve %s, in minute %d, will remain open %d\n", nextValve, 30-(time-(getDistance(graph, prev, nextValve))), time-(getDistance(graph, prev, nextValve)+1))
		time = time - (getDistance(graph, prev, nextValve) + 1)
		if time > 0 {
			p = p + (pressure[nextValve] * time)
		}
		prev = valves[i]
		i++
	}
	return p
}

func getDistance(graph map[string][]Edge, start string, end string) int {
	for _, edge := range graph[start] {
		if edge.Target == end {
			return edge.Weight
		}
	}
	return 0
}

func relevantValves(adjacency map[string][]Edge, pressure map[string]int) []string {
	keys := make([]string, 0, len(adjacency))
	for k, _ := range adjacency {
		if pressure[k] > 0 {
			keys = append(keys, k)
		}
	}
	return keys
}

func CombinationGenerator(valves []string, adjacency map[string][]Edge) chan []string {
	ch := make(chan []string)
	go func() {
		var helper func([]string, int)

		helper = func(arr []string, n int) {
			if n == 1 {
				tmp := make([]string, len(arr))
				copy(tmp, arr)
				ch <- tmp
			} else {
				for i := 0; i < n; i++ {
					helper(arr, n-1)
					if n%2 == 1 {
						tmp := arr[i]
						arr[i] = arr[n-1]
						arr[n-1] = tmp
					} else {
						tmp := arr[0]
						arr[0] = arr[n-1]
						arr[n-1] = tmp
					}
				}
			}
		}
		helper(valves, len(valves))
		close(ch)
	}()
	return ch
}

type Edge struct {
	Target string
	Weight int
}

func reduceGraph(startNode string, a map[string][]string, pressure map[string]int) map[string][]Edge {
	adjecency := make(map[string][]Edge, 0)

	for k, _ := range a {
		if pressure[k] > 0 || k == startNode {
			adjecency[k] = reduceEdges(k, k, a, pressure, 1)
		}
	}
	return adjecency
}

func reduceEdges(node string, path string, a map[string][]string, pressure map[string]int, weight int) []Edge {
	edges := make([]Edge, 0)
	for _, edge := range a[node] {
		if !strings.Contains(path, edge) {
			if pressure[edge] > 0 {
				edges = mergeEdges(edges, Edge{
					Target: edge,
					Weight: weight,
				})
			}
			edges = mergeEdges(edges, reduceEdges(edge, fmt.Sprintf("%s,%s", path, edge), a, pressure, weight+1)...)
		}
	}
	return edges
}

func mergeEdges(edges []Edge, elems ...Edge) []Edge {
	result := edges
	for _, edge := range elems {
		pos := searchEdge(result, edge)
		if pos > -1 {
			if edge.Weight < result[pos].Weight {
				result[pos].Weight = edge.Weight
			}
		} else {
			result = append(result, edge)
		}
	}
	return result
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

func searchEdge(q []Edge, u Edge) int {
	index := -1
	for i := 0; i < len(q); i++ {
		if q[i].Target == u.Target {
			index = i
		}
	}
	return index
}
