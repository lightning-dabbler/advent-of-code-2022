// https://adventofcode.com/2022/day/9

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day09-input.txt
var s string

func tailNear(tailPosition, headPosition []int) bool {
	tailX := tailPosition[0]
	tailY := tailPosition[1]
	headX := headPosition[0]
	headY := headPosition[1]
	if (headX <= tailX+1) && (headX >= tailX-1) && (headY <= tailY+1) && (headY >= tailY-1) {
		return true
	}
	return false
}

func registerPosition(positions map[string]struct{}, tailPosition []int) {
	position := fmt.Sprintf("%d,%d", tailPosition[0], tailPosition[1])
	positions[position] = struct{}{}
}

func tailMove(currentTail, previousTail []int) {
	if currentTail[0]-previousTail[0] == 0 {
		diff := previousTail[1] - currentTail[1]
		if diff > 0 {
			currentTail[1] += 1
		} else if diff < 0 {
			currentTail[1] -= 1
		}
	} else if currentTail[1]-previousTail[1] == 0 {
		diff := previousTail[0] - currentTail[0]
		if diff > 0 {
			currentTail[0] += 1
		} else if diff < 0 {
			currentTail[0] -= 1
		}
	} else {
		xdiff := previousTail[0] - currentTail[0]
		ydiff := previousTail[1] - currentTail[1]
		slope := float64(ydiff) / float64(xdiff)
		if slope > 0 {
			if xdiff < 0 && ydiff < 0 {
				currentTail[1] -= 1
				currentTail[0] -= 1
			} else {
				currentTail[1] += 1
				currentTail[0] += 1
			}
		} else {
			if xdiff < 0 {
				currentTail[0] -= 1
				currentTail[1] += 1
			} else {
				currentTail[0] += 1
				currentTail[1] -= 1
			}
		}
	}
}

func partA() {
	scanner := bufio.NewScanner(strings.NewReader(s))
	positions := map[string]struct{}{}
	tailPosition := []int{0, 0}
	headPosition := []int{0, 0}
	for scanner.Scan() {
		registerPosition(positions, tailPosition)
		move := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		direction := move[0]
		units, err := strconv.Atoi(move[1])
		if err != nil {
			log.Fatalln(err)
		}

		switch direction {
		case "R":
			for i := 1; i <= units; i++ {

				headPosition[0] += 1
				if !tailNear(tailPosition, headPosition) {
					tailMove(tailPosition, headPosition)
				}
				registerPosition(positions, tailPosition)
			}
		case "L":
			for i := 1; i <= units; i++ {
				headPosition[0] -= 1
				if !tailNear(tailPosition, headPosition) {
					tailMove(tailPosition, headPosition)
				}
				registerPosition(positions, tailPosition)

			}
		case "D":
			for i := 1; i <= units; i++ {
				headPosition[1] -= 1
				if !tailNear(tailPosition, headPosition) {
					tailMove(tailPosition, headPosition)
				}
				registerPosition(positions, tailPosition)

			}
		case "U":
			for i := 1; i <= units; i++ {
				headPosition[1] += 1
				if !tailNear(tailPosition, headPosition) {
					tailMove(tailPosition, headPosition)
				}
				registerPosition(positions, tailPosition)
			}
		default:
			log.Fatalln("Unknown " + direction)
		}
	}

	registerPosition(positions, tailPosition)
	// fmt.Println(positions)
	fmt.Println(len(positions))

}

type Position []int

func partB() {
	scanner := bufio.NewScanner(strings.NewReader(s))
	positions := map[string]struct{}{}
	tailPositions := make([]Position, 9, 9)
	for i := 0; i < 9; i++ {
		tailPositions[i] = Position{0, 0}
	}
	headPosition := []int{0, 0}

	for scanner.Scan() {
		registerPosition(positions, tailPositions[8])
		move := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		direction := move[0]
		units, err := strconv.Atoi(move[1])
		if err != nil {
			log.Fatalln(err)
		}

		switch direction {
		case "R":
			for i := 1; i <= units; i++ {
				headPosition[0] += 1
				if !tailNear(tailPositions[0], headPosition) {
					tailMove(tailPositions[0], headPosition)

					for j := 1; j < 9; j++ {
						if !tailNear(tailPositions[j], tailPositions[j-1]) {
							tailMove(tailPositions[j], tailPositions[j-1])
						} else {
							break
						}
					}
				}
				registerPosition(positions, tailPositions[8])
			}
		case "L":
			for i := 1; i <= units; i++ {
				headPosition[0] -= 1
				if !tailNear(tailPositions[0], headPosition) {
					tailMove(tailPositions[0], headPosition)

					for j := 1; j < 9; j++ {
						if !tailNear(tailPositions[j], tailPositions[j-1]) {
							tailMove(tailPositions[j], tailPositions[j-1])
						} else {
							break
						}
					}
				}
				registerPosition(positions, tailPositions[8])
			}
		case "D":
			for i := 1; i <= units; i++ {
				headPosition[1] -= 1
				if !tailNear(tailPositions[0], headPosition) {
					tailMove(tailPositions[0], headPosition)

					for j := 1; j < 9; j++ {
						if !tailNear(tailPositions[j], tailPositions[j-1]) {
							tailMove(tailPositions[j], tailPositions[j-1])
						} else {
							break
						}
					}
				}
				registerPosition(positions, tailPositions[8])

			}
		case "U":
			for i := 1; i <= units; i++ {
				headPosition[1] += 1
				if !tailNear(tailPositions[0], headPosition) {
					tailMove(tailPositions[0], headPosition)

					for j := 1; j < 9; j++ {
						if !tailNear(tailPositions[j], tailPositions[j-1]) {
							tailMove(tailPositions[j], tailPositions[j-1])
						} else {
							break
						}
					}
				}
				registerPosition(positions, tailPositions[8])
			}
		default:
			log.Fatalln("Unknown " + direction)
		}
	}

	registerPosition(positions, tailPositions[8])
	fmt.Println(len(positions))
}
func main() {
	partA() // 6098
	partB() // 2597
}
