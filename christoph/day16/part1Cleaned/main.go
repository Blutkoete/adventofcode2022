package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	//round1Brute("example.txt") // not 1752 (too high)
	round1Brute("example.txt") // not 1752 (too high), 1725 also wrong
}

func round1Brute(file string) {
	lines := readInput(file)
	adjacency, pressure := createGraph(lines)
	a2 := reduceGraph("AA", adjacency, pressure)
	combinationChannel := make(chan *Combination)
	go func() {
		maxPressure := 0
		for {
			comb := <-combinationChannel
			if comb.ReleasedPressure > maxPressure {
				fmt.Println("============================")
				fmt.Printf("New heighest combination pressure %d with path %s and released valves %s\n", comb.ReleasedPressure, strings.Join(comb.Path, ","), strings.Join(comb.ReleasedValves, ","))
				fmt.Printf("%s\n", comb.Log)
				fmt.Println("============================")
				maxPressure = comb.ReleasedPressure

			}
		}
	}()
	generateAllCombinations(combinationChannel, a2, pressure, &Combination{
		Path:             []string{"AA"},
		ReleasedValves:   []string{},
		ReleasedPressure: 0,
	}, 30)
}

func NewCombination(prev *Combination, valve string, pressure int) *Combination {
	if pressure > -1 {
		return &Combination{
			Path:             append(prev.Path, valve),
			ReleasedValves:   append(prev.ReleasedValves, valve),
			ReleasedPressure: prev.ReleasedPressure + pressure,
			Log:              prev.Log,
		}

	} else {
		return &Combination{
			Path:             append(prev.Path, valve),
			ReleasedValves:   prev.ReleasedValves,
			ReleasedPressure: prev.ReleasedPressure,
			Log:              prev.Log,
		}
	}
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

func generateAllCombinations(combinationChannel chan *Combination, adjacency map[string][]Edge, pressure map[string]int, comb *Combination, time int) {
	current := comb.Path[len(comb.Path)-1]
	if time >= 0 {
		for _, neighbor := range adjacency[current] {
			lifetimePressure := -1
			remainingTime := 0
			if !contains(comb.ReleasedValves, neighbor.Target) {
				remainingTime = time - neighbor.Weight - 1
				lifetimePressure = pressure[neighbor.Target] * remainingTime
			} else {
				remainingTime = time - neighbor.Weight
			}

			newComb := NewCombination(comb, neighbor.Target, lifetimePressure)
			newComb.Printf("Moving to %s from %s, at %d with %d remaining", neighbor.Target, current, time, remainingTime)
			combinationChannel <- newComb
			generateAllCombinations(combinationChannel, adjacency, pressure, newComb, remainingTime)
		}
	}
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

type Combination struct {
	Path             []string
	ReleasedValves   []string
	ReleasedPressure int
	Log              string
}

func (c *Combination) Println(message string) {
	c.Log = fmt.Sprintf("%s%s\n", c.Log, message)
}

func (c *Combination) Printf(format string, a ...any) {
	c.Log = fmt.Sprintf("%s%s\n", c.Log, fmt.Sprintf(format, a...))
}
