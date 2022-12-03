// https://adventofcode.com/2022/day/3

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
	"unicode"
)

//go:embed day03-input.txt
var s string

func partA() {
	scanner := bufio.NewScanner(strings.NewReader(s))

	pointsMap := map[string]int{}
	i := 1
	for r := 'a'; r <= 'z'; r++ {
		pointsMap[string(r)] = i
		pointsMap[string(unicode.ToUpper(r))] = i + 26
		i += 1
	}
	// fmt.Println(pointsMap)
	points := 0

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		firstHalfChars := map[string]struct{}{}
		// fmt.Println(text)
		n := len(text)
		// if n%2 != 0 {
		// 	fmt.Printf("Test: %s len: %d", text, n)
		// }
		for i, x := range text {
			v := string(x)
			if i+1 <= (n / 2) {
				firstHalfChars[v] = struct{}{}
			} else {
				_, ok := firstHalfChars[v]
				if ok {
					points += pointsMap[v]
					delete(firstHalfChars, v)
				}
			}
		}
	}
	fmt.Println(points)
}

func partB() {
	scanner := bufio.NewScanner(strings.NewReader(s))

	pointsMap := map[string]int{}
	i := 1
	for r := 'a'; r <= 'z'; r++ {
		pointsMap[string(r)] = i
		pointsMap[string(unicode.ToUpper(r))] = i + 26
		i += 1
	}
	// fmt.Println(pointsMap)
	points := 0

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		chars := map[string]int{}

		for _, x := range text {
			v := string(x)
			if _, ok := chars[v]; !ok {
				chars[v] += 1
			}
		}

		scanner.Scan()
		text = strings.TrimSpace(scanner.Text())

		chars2 := map[string]int{}
		for _, x := range text {
			v := string(x)
			if _, ok := chars2[v]; !ok {
				chars2[v] += 1
				chars[v] += 1
			}
		}

		scanner.Scan()
		text = strings.TrimSpace(scanner.Text())

		chars3 := map[string]int{}
		for _, x := range text {
			v := string(x)
			if _, ok := chars3[v]; !ok {
				chars3[v] += 1
				chars[v] += 1
			}
			if chars[v] == 3 {
				points += pointsMap[v]
				break
			}
		}
	}
	fmt.Println(points)
}

func main() {
	partA()
	partB()

}
