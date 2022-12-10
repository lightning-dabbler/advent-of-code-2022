// https://adventofcode.com/2022/day/10

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day10-input.txt
var s string

func partA() {
	scanner := bufio.NewScanner(strings.NewReader(s))
	cycle := 1
	// cycles := map[int]int{}
	X := 1

	nextSignalStrengthCycle := 20
	sumSignalStrengths := 0
	for scanner.Scan() {

		instruction := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(instruction) == 2 {
			value, err := strconv.Atoi(instruction[1])
			if err != nil {
				log.Fatalln(err)
			}
			for i := 1; i <= 2; i++ {
				if nextSignalStrengthCycle == cycle {
					sumSignalStrengths += (X * cycle)
					nextSignalStrengthCycle += 40
				}
				cycle += 1
			}
			X += value

		} else {
			if nextSignalStrengthCycle == cycle {
				sumSignalStrengths += (X * cycle)
				nextSignalStrengthCycle += 40
			}
			cycle += 1
		}
	}
	fmt.Println(sumSignalStrengths)
}

func partB() {
	scanner := bufio.NewScanner(strings.NewReader(s))
	type CRTRow []string
	CRT := make([]CRTRow, 6, 6)
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			CRT[i] = append(CRT[i], "#")
		}
	}
	cycle := 1
	X := 1
	for scanner.Scan() {
		instruction := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if len(instruction) == 2 {
			value, err := strconv.Atoi(instruction[1])
			if err != nil {
				log.Fatalln(err)
			}
			for i := 1; i <= 2; i++ {
				row := cycle / 40
				column := cycle % 40
				if column == 0 {
					row -= 1
					column = 39
				} else {
					column -= 1
				}
				if !(column >= X-1 && column <= X+1) {
					CRT[row][column] = " "
				}
				cycle += 1
			}
			X += value

		} else {

			row := cycle / 40
			column := cycle % 40
			if column == 0 {
				row -= 1
				column = 39
			} else {
				column -= 1
			}
			if !(column >= X-1 && column <= X+1) {
				CRT[row][column] = " "
			}
			cycle += 1
		}
	}

	for i := 0; i < 6; i++ {
		fmt.Println(strings.Join(CRT[i], ""))
	}

}
func main() {
	partA()
	partB()
}
