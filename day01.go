// https://adventofcode.com/2022/day/1
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func partA() {
	file, err := os.Open("day01-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var maxCalories int
	var total int
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			if total > maxCalories {
				maxCalories = total
			}
			total = 0
		} else {
			intVar, err := strconv.Atoi(text)
			if err != nil {
				log.Fatalln(err)
			}
			total += intVar
		}
	}
	if total > maxCalories {
		maxCalories = total
	}
	fmt.Println(maxCalories)
}

// Array Sort implementation of Part B

func partBSortArrayImplementation() {
	file, err := os.Open("day01-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var calories []int
	var total int
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			calories = append(calories, total)
			total = 0
		} else {
			intVar, err := strconv.Atoi(text)
			if err != nil {
				log.Fatalln(err)
			}
			total += intVar
		}
	}
	calories = append(calories, total)
	sort.Slice(calories, func(i, j int) bool {
		return calories[i] < calories[j]
	})
	n := len(calories)
	var sum int
	for i := 1; i <= 3; i++ {
		sum += calories[n-i]
	}
	fmt.Println(sum)
}

// Max Priority Queue Implementation of Part B

type Calories []int

func (c Calories) Len() int { return len(c) }
func (c Calories) Less(i, j int) bool {
	return c[i] > c[j]
}
func (c *Calories) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

func (c *Calories) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*c = append(*c, x.(int))
}

func (c *Calories) Peek() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	return x
}

func (c Calories) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func partBMaxHeapImplementation() {

	file, err := os.Open("day01-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	calories := &Calories{}
	heap.Init(calories)
	var total int
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			heap.Push(calories, total)
			total = 0
		} else {
			intVar, err := strconv.Atoi(text)
			if err != nil {
				log.Fatalln(err)
			}
			total += intVar
		}
	}
	heap.Push(calories, total)
	total = 0
	var sum int

	for i := 1; i <= 3; i++ {
		sum += heap.Pop(calories).(int)
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("Part A")
	partA()
	fmt.Println("Part B - Sorted Array")
	partBSortArrayImplementation()
	fmt.Println("Part B - Max Priority Queue")
	partBMaxHeapImplementation()
}
