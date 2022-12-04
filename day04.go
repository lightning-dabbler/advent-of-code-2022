// https://adventofcode.com/2022/day/4

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day04-input.txt
var s string

func partA() {
	var fullyInclusivePairCount int

	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		text := strings.TrimSpace(string(scanner.Text()))
		pair := strings.Split(text, ",")
		left, right := pair[0], pair[1]
		pair = strings.Split(left, "-")

		sectionLeft1, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatalln(err)
		}

		sectionLeft2, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatalln(err)
		}
		pair = strings.Split(right, "-")

		sectionRight1, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatalln(err)
		}
		sectionRight2, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatalln(err)
		}

		if ((sectionLeft1 <= sectionRight1) && (sectionLeft2 >= sectionRight2)) ||
			((sectionRight1 <= sectionLeft1) && (sectionRight2 >= sectionLeft2)) {
			fullyInclusivePairCount += 1
		}
	}
	fmt.Println(fullyInclusivePairCount)
}

func partB() {
	var inclusivePairCount int

	scanner := bufio.NewScanner(strings.NewReader(s))

	for scanner.Scan() {
		text := strings.TrimSpace(string(scanner.Text()))
		pair := strings.Split(text, ",")
		left, right := pair[0], pair[1]
		pair = strings.Split(left, "-")

		sectionLeft1, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatalln(err)
		}

		sectionLeft2, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatalln(err)
		}
		pair = strings.Split(right, "-")

		sectionRight1, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatalln(err)
		}
		sectionRight2, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatalln(err)
		}
		if
		// partially inclusive
		((sectionLeft1 <= sectionRight1) && (sectionLeft2 >= sectionRight1)) ||
			((sectionRight1 <= sectionLeft1) && (sectionRight2 >= sectionLeft1)) ||
			((sectionLeft1 <= sectionRight2) && (sectionLeft2 >= sectionRight2)) ||
			((sectionRight1 <= sectionLeft2) && (sectionRight2 >= sectionLeft2)) ||
			// fully inclusive
			((sectionLeft1 <= sectionRight1) && (sectionLeft2 >= sectionRight2)) ||
			((sectionRight1 <= sectionLeft1) && (sectionRight2 >= sectionLeft2)) {
			inclusivePairCount += 1
		}
	}
	fmt.Println(inclusivePairCount)
}
func main() {
	partA()
	partB()
}
