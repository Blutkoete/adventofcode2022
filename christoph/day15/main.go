package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Printf("example.txt %f\n", round1("example.txt", 10))
	fmt.Printf("input.txt %f\n", round1("input.txt", 2000000))
	fmt.Printf("example.txt %f\n", round2("example.txt", 10))
	fmt.Printf("input.txt %f\n", round2("input.txt", 2000000))
}

var min = 0
var max = 4000000

type Sensor struct {
	atX     int
	atY     int
	beaconX int
	beaconY int
	dist    float64
}

func round1(file string, rowId int) float64 {
	lines := readInput(file)
	sensors := parseInput(lines)
	intervals := getIntervals(sensors, rowId)
	regions := make([][]float64, 0)
	sum := 0.0
	for _, interval := range intervals {
		sum = sum + (volume(interval) - intersectionWithRegionsVolume(interval, regions))
		regions = append(regions, interval)
	}
	return sum
}

func round2(file string, rowId int) float64 {
	lines := readInput(file)
	sensors := parseInput(lines)
	for blindsOnRow(sensors, rowId) == 0 {
		rowId++
	}
	intervals := getIntervals(sensors, rowId)
	for x := min; x <= max; x++ {
		if allIntersect(float64(x), intervals) {
			return float64(x)*float64(max) + float64(rowId)
		}
	}
	regions := make([][]float64, 0)
	sum := 0.0
	for _, interval := range intervals {
		sum = sum + (volume(interval) - intersectionWithRegionsVolume(interval, regions))
		regions = append(regions, interval)
	}
	return sum
}

func allIntersect(x float64, intervals [][]float64) bool {
	for _, i := range intervals {
		if !(i[0] > x || i[1] < x) {
			return false
		}
	}
	return true
}

func parseInput(lines []string) []*Sensor {
	sensors := make([]*Sensor, 0)
	for _, line := range lines {
		re := regexp.MustCompile(".*x=([-0-9]*).*y=([-0-9]*).*x=([-0-9]*).*y=([-0-9]*)")
		matches := re.FindStringSubmatch(line)
		sensorAtX, err := strconv.Atoi(matches[1])
		checkError(err)
		sensorAtY, err := strconv.Atoi(matches[2])
		checkError(err)
		beaconAtX, err := strconv.Atoi(matches[3])
		checkError(err)
		beaconAtY, err := strconv.Atoi(matches[4])
		checkError(err)
		sensors = append(sensors, &Sensor{
			atX:     sensorAtX,
			atY:     sensorAtY,
			beaconX: beaconAtX,
			beaconY: beaconAtY,
			dist:    dist(sensorAtX, sensorAtY, beaconAtX, beaconAtY),
		})
	}

	return sensors
}

func dist(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2))
}

func volume(bounds []float64) float64 {
	return math.Abs(bounds[1] - bounds[0])
}

func regionIntersect(region1 []float64, region2 []float64) bool {
	return region2[0] <= region1[1] && region2[1] >= region1[0]
}

func blindsOnRow(sensors []*Sensor, y int) float64 {
	return float64(max) - intersectionWithRegionsVolume([]float64{float64(min), float64(max)}, getIntervals(sensors, y))
}

func getIntervals(sensors []*Sensor, y int) [][]float64 {
	intervals := make([][]float64, 0)
	for _, sensor := range sensors {
		spareX := sensor.dist - math.Abs(float64(sensor.atY-y))
		if spareX >= 0 {
			intervals = append(intervals, []float64{float64(sensor.atX) - spareX, float64(sensor.atX) + spareX})
		}
	}
	return intervals
}

func intersectionWithRegionsVolume(reg []float64, regions [][]float64) float64 {
	x := 0.0
	for i, region := range regions {
		if !regionIntersect(reg, region) {
			continue
		}
		tmpBounds := []float64{math.Max(region[0], reg[0]), math.Min(region[1], reg[1])}
		x = x + (volume(tmpBounds) - intersectionWithRegionsVolume(tmpBounds, regions[i+1:]))

	}
	return x
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
