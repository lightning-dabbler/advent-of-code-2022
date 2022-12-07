// https://adventofcode.com/2022/day/6

package main

import (
	_ "embed"
	"fmt"
)

//go:embed day06-input.txt
var s string

func partA() {
	characterIndex := 1
	n := len(s)
	for i := 0; i < n-4; i++ {
		v := s[i]
		unique := true
		characterIndex = i + 1
		Map := map[byte]struct{}{v: {}}
		for j := i + 1; j < i+4; j++ {
			_, ok := Map[s[j]]
			if ok {
				unique = false
				break
			} else {
				Map[s[j]] = struct{}{}
				characterIndex += 1
			}
		}
		if unique {
			break
		}
	}
	fmt.Println(characterIndex)
}

func partB() {
	characterIndex := 1
	n := len(s)
	for i := 0; i < n-14; i++ {
		v := s[i]
		unique := true
		characterIndex = i + 1
		Map := map[byte]struct{}{v: {}}
		for j := i + 1; j < i+14; j++ {
			_, ok := Map[s[j]]
			if ok {
				unique = false
				break
			} else {
				Map[s[j]] = struct{}{}
				characterIndex += 1
			}
		}
		if unique {
			break
		}
	}
	fmt.Println(characterIndex)
}

func main() {
	partA()
	partB()
}
