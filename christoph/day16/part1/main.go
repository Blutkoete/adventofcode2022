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
	//round1Brute("example.txt") // not 1752 (too high)
	round1Brute("input.txt") // not 1752 (too high), 1725 also wrong
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

func NewCombination(prev *Combination, valve string, released bool, pressure int) *Combination {
	if released {
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
			if !contains(comb.ReleasedValves, neighbor.Target) {
				remainingTime := time - neighbor.Weight - 1
				lifetimePressure := pressure[neighbor.Target] * remainingTime
				//comb.Printf("    Option %s with lifetimePressure %d", neighbor.Target, lifetimePressure)
				if lifetimePressure >= 0 {
					newComb := NewCombination(comb, neighbor.Target, true, lifetimePressure)
					newComb.Printf("Moving to %s from %s at %d, to open valve %d lifetimePressure %d and remaining %d", neighbor.Target, current, time, pressure[neighbor.Target], lifetimePressure, remainingTime)
					combinationChannel <- newComb
					generateAllCombinations(combinationChannel, adjacency, pressure, newComb, remainingTime)
				}
			} else {
				remainingTime := time - neighbor.Weight
				newComb := NewCombination(comb, neighbor.Target, false, 0)
				newComb.Printf("Moving to %s from %s, at %d with %d remaining", neighbor.Target, current, time, remainingTime)
				combinationChannel <- newComb
				generateAllCombinations(combinationChannel, adjacency, pressure, newComb, remainingTime)
			}
		}
	}
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

func searchEdge(q []Edge, u Edge) int {
	index := -1
	for i := 0; i < len(q); i++ {
		if q[i].Target == u.Target {
			index = i
		}
	}
	return index
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

/*

Maybe future optimization
// Get all distances https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm
func floydWarshall(adjacency map[string][]string) {
	dist := make(map[string])

}


let dist be a |V| × |V| array of minimum distances initialized to ∞ (infinity)
for each edge (u, v) do
    dist[u][v] ← w(u, v)  // The weight of the edge (u, v)
for each vertex v do
    dist[v][v] ← 0
for k from 1 to |V|
    for i from 1 to |V|
        for j from 1 to |V|
            if dist[i][j] > dist[i][k] + dist[k][j]
                dist[i][j] ← dist[i][k] + dist[k][j]
            end if
*/
