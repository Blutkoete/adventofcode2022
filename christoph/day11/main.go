package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"math/big"

	progressbar "github.com/schollz/progressbar/v3"
)

var zero = big.NewInt(0)

type Monkey struct {
	Index            int
	Items            []*big.Int
	Operation        string
	OperationValue   *big.Int
	OperationWithOld bool
	TestDivisor      *big.Int
	TestTrueMonkey   int
	TestFalseMonkey  int
	InspectionCount  int64
}

func main() {
	calculate("example.txt", 20, true)
	calculate("input.txt", 20, true)
	//calculate("example.txt", 10000, false)
	calculate("input.txt", 10000, false)
}

func calculate(file string, iterations int, divideByThree bool) {
	fmt.Printf("Calculating Monkey Business for %s %d\n", file, iterations)
	lines := readInput(file)
	monkeys, bcd := parseMonkeys(lines)
	bar := progressbar.Default(int64(iterations))
	for i := 0; i < iterations; i++ {
		//fmt.Printf("Iteration %d\n", i)
		bar.Add(1)
		for _, monkey := range monkeys {
			//fmt.Printf("Monkey %d:\n", monkey.Index)
			for _, item := range monkey.Items {
				monkey.InspectionCount++
				//fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", item)
				inspection := updateWorry(item, monkey)
				if divideByThree {
					inspection = inspection.Div(inspection, big.NewInt(3))
				} else if inspection.Cmp(big.NewInt(bcd)) >= 1 {
					inspection = inspection.Mod(inspection, big.NewInt(bcd))
				}

				//fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", inspection)
				modulus := big.NewInt(0)
				modulus.Mod(inspection, monkey.TestDivisor)
				if modulus.Uint64() == 0 {
					//fmt.Printf("    Current worry level is divisible by %d.\n", monkey.TestDivisor)
					//fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", inspection, monkey.TestTrueMonkey)
					monkeys[monkey.TestTrueMonkey].Items = append(monkeys[monkey.TestTrueMonkey].Items, inspection)
				} else {
					//fmt.Printf("    Current worry level is not divisible by %d.\n", monkey.TestDivisor)
					//fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", inspection, monkey.TestFalseMonkey)
					monkeys[monkey.TestFalseMonkey].Items = append(monkeys[monkey.TestFalseMonkey].Items, inspection)
				}
			}
			monkey.Items = make([]*big.Int, 0)
		}
	}
	values := make([]int, 0)
	for _, monkey := range monkeys {
		//fmt.Printf("Monkey %d insepctions %d\n", monkey.Index, monkey.InspectionCount)
		values = append(values, int(monkey.InspectionCount))
	}
	sort.Ints(values)

	fmt.Printf("Monkey Business for %s %d %d\n", file, iterations, int64(values[len(values)-1]*values[len(values)-2]))
	fmt.Println("-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-")
}

func updateWorry(initialConcern *big.Int, monkey *Monkey) *big.Int {
	if monkey.Operation == "+" {
		if monkey.OperationWithOld {
			initialConcern.Add(initialConcern, initialConcern)
			//fmt.Printf("    Worry level increases by itself to %d.\n", result)
			return initialConcern
		} else {
			initialConcern.Add(initialConcern, monkey.OperationValue)
			//fmt.Printf("    Worry level increases by %d to %d.\n", monkey.OperationValue, result)
			return initialConcern
		}
	} else if monkey.Operation == "*" {
		if monkey.OperationWithOld {
			//fmt.Printf("    Worry level is multiplied by itself to %d.\n", result)
			return initialConcern.Mul(initialConcern, initialConcern)
		} else {
			//fmt.Printf("    Worry level is multiplied by %d to %d.\n", monkey.OperationValue, result)
			return initialConcern.Mul(initialConcern, monkey.OperationValue)
		}
	} else {
		log.Panicf("Unknown operator %s", monkey.Operation)
		return nil
	}
}

func parseMonkeys(lines []string) ([]*Monkey, int64) {
	monkeys := make([]*Monkey, 0)
	bcd := 1
	var currentMonkey *Monkey
	for _, line := range lines {
		if strings.HasPrefix(line, "Monkey") {
			indexStr := line[7:8]
			index, err := strconv.Atoi(indexStr)
			checkError(err)

			currentMonkey = &Monkey{
				Index:            index,
				InspectionCount:  0,
				Operation:        "",
				OperationValue:   big.NewInt(0),
				OperationWithOld: false,
				TestDivisor:      big.NewInt(0),
				TestTrueMonkey:   0,
				TestFalseMonkey:  0,
			}
			monkeys = append(monkeys, currentMonkey)
		} else if strings.HasPrefix(line, "  Starting items: ") {
			itemsStr := line[17:]
			itemsArray := strings.Split(itemsStr, ", ")
			items := make([]*big.Int, 0)
			for _, itemStr := range itemsArray {
				itemConcern, err := strconv.Atoi(strings.TrimSpace(itemStr))
				checkError(err)
				items = append(items, big.NewInt(int64(itemConcern)))
			}
			currentMonkey.Items = items
		} else if strings.HasPrefix(line, "  Operation: new = old ") {
			ops := line[23:]
			opsParts := strings.Split(ops, " ")
			currentMonkey.Operation = strings.TrimSpace(opsParts[0])
			if opsParts[1] != "old" {
				number, err := strconv.Atoi(strings.TrimSpace(opsParts[1]))
				checkError(err)
				currentMonkey.OperationValue = big.NewInt(int64(number))
				currentMonkey.OperationWithOld = false
			} else {
				currentMonkey.OperationWithOld = true
			}
		} else if strings.HasPrefix(line, "  Test: divisible by ") {
			divisorStr := line[21:]
			divisor, err := strconv.Atoi(strings.TrimSpace(divisorStr))
			checkError(err)
			bcd = bcd * divisor
			currentMonkey.TestDivisor = big.NewInt(int64(divisor))
		} else if strings.HasPrefix(line, "    If true: throw to monkey ") {
			numberStr := line[29:]
			number, err := strconv.Atoi(strings.TrimSpace(numberStr))
			checkError(err)
			currentMonkey.TestTrueMonkey = number
		} else if strings.HasPrefix(line, "    If false: throw to monkey ") {
			numberStr := line[29:]
			number, err := strconv.Atoi(strings.TrimSpace(numberStr))
			checkError(err)
			currentMonkey.TestFalseMonkey = number
		} else if line != "" {
			fmt.Printf("Unrecognized command '%s'\n", line)
		}
	}
	return monkeys, int64(bcd)
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
