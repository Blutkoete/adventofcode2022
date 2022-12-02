package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	round2()
}

func round1() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	roundScore := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		s := strings.Split(line, " ")
		selection := s[0]
		strategy := s[1]
		fmt.Printf("Round %s %s\n",selection,strategy)
		if strategy == "X" { // I play rock
			roundScore = roundScore + 1
			if selection == "A" { //rock
				roundScore = roundScore + 3 //tie
			} else if selection == "B" { // paper
				roundScore = roundScore + 0 //lost
			} else if selection == "C" { // sicssors
				roundScore = roundScore + 6 //won 
			}
		} else if strategy == "Y" { // I play Paper
			roundScore = roundScore + 2
			if selection == "A" { //rock
				roundScore = roundScore + 6 //won
			} else if selection == "B" { // paper
				roundScore = roundScore + 3 //tie
			} else if selection == "C" { // sicssors
				roundScore = roundScore + 0 //lost 
			}
		} else if strategy == "Z" {// I play scissors
			roundScore = roundScore + 3
			if selection == "A" { //rock
				roundScore = roundScore + 0 //lost
			} else if selection == "B" { // paper
				roundScore = roundScore + 6 //won
			} else if selection == "C" { // sicssors
				roundScore = roundScore + 3 //tie 
			}
			
		}
	}
	fmt.Printf("Score %d\n",roundScore)
}


func round2() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	roundScore := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		s := strings.Split(line, " ")
		selection := s[0]
		strategy := s[1]
		fmt.Printf("Round %s %s\n",selection,strategy)
		if strategy == "X" { // lose
			roundScore = roundScore + 0
			if selection == "A" { //rock --> scissors
				roundScore = roundScore + 3
			} else if selection == "B" { // paper --> rock
				roundScore = roundScore +1
			} else if selection == "C" { // sicssors --> paper
				roundScore = roundScore + 2
			}
		} else if strategy == "Y" { // draw
			roundScore = roundScore + 3
			if selection == "A" { //rock --> rock
				roundScore = roundScore + 1
			} else if selection == "B" { // paper --> paper
				roundScore = roundScore + 2
			} else if selection == "C" { // sicssors --> sicssors
				roundScore = roundScore + 3
			}
		} else if strategy == "Z" {// win
			roundScore = roundScore + 6
			if selection == "A" { //rock --> paper
				roundScore = roundScore + 2
			} else if selection == "B" { // paper --> scissors
				roundScore = roundScore +3
			} else if selection == "C" { // sicssors --> rock
				roundScore = roundScore + 1
			}
		}
	}
	fmt.Printf("Score %d\n",roundScore)
}