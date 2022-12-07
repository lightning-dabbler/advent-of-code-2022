// https://adventofcode.com/2022/day/7

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day07-input.txt
var s string

func partA() {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var cwd string
	var cwdStruct []string

	dirSizes := map[string]int{}
	filesFound := map[string]struct{}{}

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		console := strings.Split(text, " ")

		if console[0] == "$" {
			curCommand := console[1]
			if curCommand == "cd" {
				path := console[2]
				if path == ".." {
					cwdStruct = cwdStruct[:len(cwdStruct)-1]
					cwd = "/" + strings.TrimLeft(strings.Join(cwdStruct, "/"), "/")
				} else if path == "/" {
					cwdStruct = []string{"/"}
					cwd = "/"
				} else {
					trimmedString := strings.Trim(path, "/")
					paths := strings.Split(trimmedString, "/")
					cwdStruct = append(cwdStruct, paths...)
					cwd = strings.TrimRight(cwd, "/") + "/" + trimmedString
				}
			}

		} else if console[0] != "dir" {
			size, err := strconv.Atoi(console[0])
			if err != nil {
				log.Fatalln(err)
			}
			file := strings.TrimRight(cwd, "/") + "/" + console[1]
			if _, ok := filesFound[file]; !ok {
				// dirSizes[cwd] += size
				n := len(cwdStruct)
				for i := n - 1; i >= 0; i-- {
					dir := "/" + strings.TrimLeft(strings.Join(cwdStruct[:i+1], "/"), "/")
					dirSizes[dir] += size
				}
				filesFound[file] = struct{}{}
			}

		}
	}
	total := 0
	for _, v := range dirSizes {
		if v <= 100000 {
			total += v
		}
	}
	fmt.Println(total)
}

func partB() {
	scanner := bufio.NewScanner(strings.NewReader(s))
	var cwd string
	var cwdStruct []string

	dirSizes := map[string]int{}
	filesFound := map[string]struct{}{}
	diskSpaceSize := 70000000
	updateSize := 30000000
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		console := strings.Split(text, " ")

		if console[0] == "$" {
			curCommand := console[1]
			if curCommand == "cd" {
				path := console[2]
				if path == ".." {
					cwdStruct = cwdStruct[:len(cwdStruct)-1]
					cwd = "/" + strings.TrimLeft(strings.Join(cwdStruct, "/"), "/")
				} else if path == "/" {
					cwdStruct = []string{"/"}
					cwd = "/"
				} else {
					trimmedString := strings.Trim(path, "/")
					paths := strings.Split(trimmedString, "/")
					cwdStruct = append(cwdStruct, paths...)
					cwd = strings.TrimRight(cwd, "/") + "/" + trimmedString
				}
			}

		} else if console[0] != "dir" {
			size, err := strconv.Atoi(console[0])
			if err != nil {
				log.Fatalln(err)
			}
			file := strings.TrimRight(cwd, "/") + "/" + console[1]
			if _, ok := filesFound[file]; !ok {
				// dirSizes[cwd] += size
				n := len(cwdStruct)
				for i := n - 1; i >= 0; i-- {
					dir := "/" + strings.TrimLeft(strings.Join(cwdStruct[:i+1], "/"), "/")
					dirSizes[dir] += size
				}
				filesFound[file] = struct{}{}
			}

		}
	}

	SpaceNeededToUpdate := diskSpaceSize - dirSizes["/"]

	if SpaceNeededToUpdate > updateSize {
		panic("No deletion needed")
	} else {
		diff := updateSize - SpaceNeededToUpdate
		optimalSize := 0
		for _, v := range dirSizes {
			if optimalSize == 0 {
				if v >= diff {
					optimalSize = v
				}
			} else {
				if (optimalSize > v) && (v >= diff) {
					optimalSize = v
				}
			}
			if optimalSize == diff {
				break
			}
		}
		fmt.Println(optimalSize)
	}

}

func main() {

	partA()
	partB()
}
