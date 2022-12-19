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
	"time"
)

var wg sync.WaitGroup
var results []int
var possibleScoreByTimeLeft []int

func main() {
	possibleScoreByTimeLeft = generatePossibleScoreByTimeLeft(40)
	//solve("input.txt", 24, false)
	solve("input.txt", 32, true)
}

type Result struct {
	BlueprintId int
	Geode       int
}

func solve(file string, runtime int, topThree bool) {
	lines := readInput(file)
	blueprints := parseBlueprints(lines)
	results = make([]int, len(blueprints)+1)
	results[1] = 40
	mu := sync.Mutex{}
	channel := make(chan Result)
	go func() {
		for result := range channel {
			mu.Lock()
			if result.Geode > results[result.BlueprintId] {
				results[result.BlueprintId] = result.Geode
				fmt.Printf("Blueprint %d generates %d geodes\n", result.BlueprintId, result.Geode)
			}
			mu.Unlock()
		}
	}()
	if topThree {
		blueprints = blueprints[0:3]
	}
	for _, blueprint := range blueprints {
		if blueprint.Id == 1 {
			continue
		}
		fmt.Printf("Simulating blueprint %d\n", blueprint.Id)
		storage := Resource{
			Ore:      0,
			Clay:     0,
			Obsidian: 0,
			Geode:    0,
		}
		robots := []int{1, 0, 0, 0}
		go permutate(channel, blueprint, robots, storage, runtime)
	}
	time.Sleep(1 * time.Second)
	wg.Wait()
	qualityLevel := 0
	for _, blueprint := range blueprints {
		qualityLevel = qualityLevel + (blueprint.Id * results[blueprint.Id])
	}
	fmt.Printf("Total Quality Level is %d \n", qualityLevel)

	multi := 1
	for _, blueprint := range blueprints {
		multi = multi * results[blueprint.Id]
	}
	fmt.Printf("Multiplied result %d \n", multi)
}

func permutate(channel chan Result, blueprint Blueprint, robots []int, storage Resource, runtime int) {
	wg.Add(1)
	defer wg.Done()
	if runtime == 0 {
		if storage.Geode > 0 {
			channel <- Result{
				BlueprintId: blueprint.Id,
				Geode:       storage.Geode,
			}
		}
		return
	}
	currentGeodes := storage.Geode
	addedGeodesFromCurrentRobots := robots[3] * (runtime + 1)
	maxAddedGeodesFromFutureRobots := possibleScoreByTimeLeft[runtime]
	// prune branches that can't yield more geodes then the current max (assuming we are adding a robot every minute)
	if currentGeodes+addedGeodesFromCurrentRobots+maxAddedGeodesFromFutureRobots < results[blueprint.Id]-1 {
		return
	} else {
		oreRobotPossible, clayRobotPossible, obsidianRobotPossible, geodeRobotPossible, allowSkipping := getPossibleActions(blueprint, storage, robots)
		storage = harvest(robots, storage)
		if geodeRobotPossible {
			newRobots := []int{0, 0, 0, 0}
			copy(newRobots, robots)
			newRobots[3]++
			permutate(channel, blueprint, newRobots, reduceByCost(storage, blueprint.GeodeRobotCost), runtime-1)
		}
		if obsidianRobotPossible {
			newRobots := []int{0, 0, 0, 0}
			copy(newRobots, robots)
			newRobots[2]++
			permutate(channel, blueprint, newRobots, reduceByCost(storage, blueprint.ObsidianRobotCost), runtime-1)
		}
		if clayRobotPossible {
			newRobots := []int{0, 0, 0, 0}
			copy(newRobots, robots)
			newRobots[1]++
			permutate(channel, blueprint, newRobots, reduceByCost(storage, blueprint.ClayRobotCost), runtime-1)
		}
		if oreRobotPossible {
			newRobots := []int{0, 0, 0, 0}
			copy(newRobots, robots)
			newRobots[0]++
			permutate(channel, blueprint, newRobots, reduceByCost(storage, blueprint.OreRobotCost), runtime-1)
		}
		if allowSkipping {
			permutate(channel, blueprint, robots, storage, runtime-1)
		}
	}
}

func reduceByCost(storage Resource, cost Resource) Resource {
	newStorage := storage
	newStorage.Ore = newStorage.Ore - cost.Ore
	newStorage.Clay = newStorage.Clay - cost.Clay
	newStorage.Obsidian = newStorage.Obsidian - cost.Obsidian
	newStorage.Geode = newStorage.Geode - cost.Geode

	return newStorage
}

func harvest(robots []int, storage Resource) Resource {
	storage.Ore = storage.Ore + robots[0]
	storage.Clay = storage.Clay + robots[1]
	storage.Obsidian = storage.Obsidian + robots[2]
	storage.Geode = storage.Geode + robots[3]
	return storage
}

func max(values ...int) int {
	maxVal := -1
	for _, val := range values {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func getPossibleActions(blueprint Blueprint, storage Resource, robots []int) (bool, bool, bool, bool, bool) {
	oreRobot := canBuildRobot(storage, blueprint.OreRobotCost)
	if robots[0] >= max(blueprint.OreRobotCost.Ore, blueprint.ClayRobotCost.Ore, blueprint.ObsidianRobotCost.Ore, blueprint.GeodeRobotCost.Ore) {
		oreRobot = false
	}
	clayRobot := canBuildRobot(storage, blueprint.ClayRobotCost)
	if robots[1] >= max(blueprint.OreRobotCost.Clay, blueprint.ClayRobotCost.Clay, blueprint.ObsidianRobotCost.Clay, blueprint.GeodeRobotCost.Clay) {
		clayRobot = false
	}
	obsidianRobot := canBuildRobot(storage, blueprint.ObsidianRobotCost)
	if robots[2] >= max(blueprint.OreRobotCost.Obsidian, blueprint.ClayRobotCost.Obsidian, blueprint.ObsidianRobotCost.Obsidian, blueprint.GeodeRobotCost.Obsidian) {
		obsidianRobot = false
	}
	geodeRobot := canBuildRobot(storage, blueprint.GeodeRobotCost)
	allowSkipping := true
	if robots[0] >= max(blueprint.OreRobotCost.Ore, blueprint.ClayRobotCost.Ore, blueprint.ObsidianRobotCost.Ore, blueprint.GeodeRobotCost.Ore) && robots[1] >= max(blueprint.OreRobotCost.Clay, blueprint.ClayRobotCost.Clay, blueprint.ObsidianRobotCost.Clay, blueprint.GeodeRobotCost.Clay) && robots[2] >= max(blueprint.OreRobotCost.Obsidian, blueprint.ClayRobotCost.Obsidian, blueprint.ObsidianRobotCost.Obsidian, blueprint.GeodeRobotCost.Obsidian) {
		allowSkipping = false
	}
	return oreRobot, clayRobot, obsidianRobot, geodeRobot, allowSkipping
}

func canBuildRobot(storage Resource, cost Resource) bool {
	if storage.Ore >= cost.Ore {
		if storage.Clay >= cost.Clay {
			if storage.Obsidian >= cost.Obsidian {
				return true
			}
		}
	}
	return false
}

type Resource struct {
	Ore      int
	Clay     int
	Obsidian int
	Geode    int
}

type Blueprint struct {
	Id                int
	OreRobotCost      Resource
	ClayRobotCost     Resource
	ObsidianRobotCost Resource
	GeodeRobotCost    Resource
}

func parseBlueprints(lines []string) []Blueprint {
	blueprints := make([]Blueprint, 0)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		id := getBlueprintId(parts[0])
		robotParts := strings.Split(parts[1], ".")
		blueprint := Blueprint{
			Id: id,
		}

		for _, robotPart := range robotParts {
			if strings.HasPrefix(robotPart, " Each ore") {
				blueprint.OreRobotCost = Resource{
					Ore: getOreCost(robotPart),
				}
			} else if strings.HasPrefix(robotPart, " Each clay") {
				blueprint.ClayRobotCost = Resource{
					Ore: getOreCost(robotPart),
				}
			} else if strings.HasPrefix(robotPart, " Each obsidian") {
				ore, clay := getOreAndClayCost(robotPart)
				blueprint.ObsidianRobotCost = Resource{
					Ore:  ore,
					Clay: clay,
				}
			} else if strings.HasPrefix(robotPart, " Each geode") {
				ore, obsidian := getOreAndObsidianCost(robotPart)
				blueprint.GeodeRobotCost = Resource{
					Ore:      ore,
					Obsidian: obsidian,
				}
			}
		}

		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

func getOreAndClayCost(part string) (int, int) {
	re := regexp.MustCompile("Each (.*) robot costs ([0-9]*) ore and ([0-9]*) clay")
	matches := re.FindStringSubmatch(strings.TrimSpace(part))
	oreCost, err := strconv.Atoi(matches[2])
	checkError(err)
	clayCost, err := strconv.Atoi(matches[3])
	checkError(err)
	return oreCost, clayCost
}

func getOreAndObsidianCost(part string) (int, int) {
	re := regexp.MustCompile("Each (.*) robot costs ([0-9]*) ore and ([0-9]*) obsidian")
	matches := re.FindStringSubmatch(strings.TrimSpace(part))
	oreCost, err := strconv.Atoi(matches[2])
	checkError(err)
	obsidianCost, err := strconv.Atoi(matches[3])
	checkError(err)
	return oreCost, obsidianCost
}

func getOreCost(part string) int {
	re := regexp.MustCompile("Each (.*) robot costs ([0-9]*) ore")
	matches := re.FindStringSubmatch(strings.TrimSpace(part))
	oreCost, err := strconv.Atoi(matches[2])
	checkError(err)
	return oreCost
}

func getBlueprintId(part string) int {
	re := regexp.MustCompile("Blueprint ([0-9]*)")
	matches := re.FindStringSubmatch(part)
	blueprintId, err := strconv.Atoi(matches[1])
	checkError(err)
	return blueprintId
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

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func generatePossibleScoreByTimeLeft(maxTime int) []int {
	result := make([]int, maxTime)
	for i := 0; i < maxTime; i++ {
		temp := 0
		for j := 1; j <= i; j++ {
			temp = temp + j
		}
		result[i] = temp
	}
	return result
}
