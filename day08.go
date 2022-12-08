// https://adventofcode.com/2022/day/8

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day08-input.txt
var s string

func partA() {
	scanner := bufio.NewScanner(strings.NewReader(s))
	grid := [][]int{}
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		row := make([]int, len(text), len(text))
		for i, v := range text {
			tree, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatalln(err)
			}
			row[i] = tree
		}
		grid = append(grid, row)
	}
	// fmt.Println(grid)
	n := len(grid)
	m := len(grid[0])
	interiorVisibleCount := 0
	for i := 1; i < m-1; i++ {
	outer:
		for j := 1; j < n-1; j++ {
			tree := grid[j][i]
			leftMax := 0
			rightMax := 0
			upMax := 0
			downMax := 0
			for k := 0; k < i; k++ {
				if grid[j][k] > leftMax {
					leftMax = grid[j][k]
				}
			}
			for k := m - 1; k > i; k-- {
				if grid[j][k] > rightMax {
					rightMax = grid[j][k]
				}
			}
			for k := 0; k < j; k++ {
				if grid[k][i] > upMax {
					upMax = grid[k][i]
				}
			}
			for k := n - 1; k > j; k-- {
				if grid[k][i] > downMax {
					downMax = grid[k][i]
				}
			}

			if leftMax < tree {
				interiorVisibleCount += 1
				continue outer
			}
			if rightMax < tree {
				interiorVisibleCount += 1
				continue outer
			}
			if upMax < tree {
				interiorVisibleCount += 1
				continue outer
			}
			if downMax < tree {
				interiorVisibleCount += 1
				continue outer
			}
		}
	}
	fmt.Println(interiorVisibleCount + (2*n + 2*(m-2)))
}

func partB() {
	scanner := bufio.NewScanner(strings.NewReader(s))
	grid := [][]int{}
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		row := make([]int, len(text), len(text))
		for i, v := range text {
			tree, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatalln(err)
			}
			row[i] = tree
		}
		grid = append(grid, row)
	}
	// fmt.Println(grid)
	n := len(grid)
	m := len(grid[0])
	highestScenicScore := 0

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			tree := grid[j][i]
			distanceToLeftEdge := 0
			distanceToRightEdge := 0
			distanceToUpEdge := 0
			distanceToDownEdge := 0
			for k := i - 1; k >= 0; k-- {
				distanceToLeftEdge += 1
				if grid[j][k] >= tree {
					break
				}
			}
			for k := i + 1; k < m; k++ {
				distanceToRightEdge += 1
				if grid[j][k] >= tree {
					break
				}
			}
			for k := j - 1; k >= 0; k-- {
				distanceToUpEdge += 1
				if grid[k][i] >= tree {
					break
				}
			}
			for k := j + 1; k < n; k++ {
				distanceToDownEdge += 1
				if grid[k][i] >= tree {
					break
				}
			}
			scenicScore := distanceToDownEdge * distanceToUpEdge * distanceToRightEdge * distanceToLeftEdge
			// fmt.Println(scenicScore, j, i, tree)
			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}

		}
	}
	fmt.Println(highestScenicScore)
}

func main() {
	partA()
	partB()
}
