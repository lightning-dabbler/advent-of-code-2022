// https://adventofcode.com/2022/day/2

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func partA() {
	file, err := os.Open("day02-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	characterMapping := map[string]string{
		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	enemyMapping := map[string]string{
		"X": "C",
		"Y": "A",
		"Z": "B",
	}

	pointsMapping := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	totalScore := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		round := strings.Split(text, " ")
		var decision string
		opponentMove, myMove := round[0], round[1]
		if characterMapping[myMove] == opponentMove {
			decision = "draw"
		} else if enemyMapping[myMove] == opponentMove {
			decision = "win"
		}

		if decision == "draw" {
			totalScore += 3
		} else if decision == "win" {
			totalScore += 6
		}
		totalScore += pointsMapping[myMove]
	}
	fmt.Println(totalScore)
}

func partB() {
	file, err := os.Open("day02-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	PreditorToPreyMapping := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}

	PreyToPreditorMapping := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}

	pointsMapping := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	totalScore := 0
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		round := strings.Split(text, " ")
		opponentMove, myMove := round[0], round[1]
		if myMove == "X" {
			totalScore += pointsMapping[PreditorToPreyMapping[opponentMove]]
		} else if myMove == "Y" {
			totalScore += pointsMapping[opponentMove]
			totalScore += 3
		} else {
			totalScore += pointsMapping[PreyToPreditorMapping[opponentMove]]
			totalScore += 6
		}
	}
	fmt.Println(totalScore)
}

func main() {
	partA()
	partB()
}
