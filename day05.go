package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day05-input.txt
var s string

type Queue struct {
	Crates []string
	Size   int
}

func (s *Queue) Pop() string {
	value := s.Crates[0]
	s.Crates = s.Crates[1:]
	s.Size -= 1
	return value
}

func (s *Queue) Append(crate string) {
	s.Crates = append(s.Crates, crate)
	s.Size += 1
}

func (s *Queue) AppendLeft(crate string) {
	s.Crates = append(s.Crates, "")
	copy(s.Crates[1:], s.Crates)

	s.Crates[0] = crate
	s.Size += 1
}

func (s *Queue) Peak() string {
	return s.Crates[0]
}

func partA() {

	scanner := bufio.NewScanner(strings.NewReader(s))

	var stacks []Queue
	firstIteration := true
	for scanner.Scan() {
		text := scanner.Text()
		n := len(text)
		m := (n + 1) / 4
		if firstIteration {
			stacks = make([]Queue, m, m)
			firstIteration = false
			index := 1
			for i := 0; i < m; i++ {
				if (index < n) && strings.TrimSpace(string(text[index])) != "" {
					stacks[i].Append(string(text[index]))
				}
				index += 4
			}
			var numbersFound bool
			for !numbersFound {
				scanner.Scan()
				text = scanner.Text()
				if string(text[1]) == "1" {
					break
				}
				index = 1
				for i := 0; i < m; i++ {
					if (index < n) && strings.TrimSpace(string(text[index])) != "" {
						stacks[i].Append(string(text[index]))
					}
					index += 4
				}
			}
		} else {
			text = strings.TrimSpace(text)
			if text != "" {
				statement := strings.Split(text, " ")
				nMoves, err := strconv.Atoi(statement[1])
				if err != nil {
					log.Fatalln(err)
				}
				initialQueue, err := strconv.Atoi(statement[3])
				if err != nil {
					log.Fatalln(err)
				}
				finalQueue, err := strconv.Atoi(statement[5])
				if err != nil {
					log.Fatalln(err)
				}
				var cratesToTransport []string
				for i := 0; i < nMoves; i++ {
					crate := stacks[initialQueue-1].Pop()
					cratesToTransport = append(cratesToTransport, crate)
				}

				for i := 0; i < nMoves; i++ {
					crate := cratesToTransport[i]
					stacks[finalQueue-1].AppendLeft(crate)
				}
			}
		}
	}
	var output string
	for _, v := range stacks {
		output += v.Peak()
	}
	fmt.Println(output)
}

func partB() {

	scanner := bufio.NewScanner(strings.NewReader(s))

	var stacks []Queue
	firstIteration := true
	for scanner.Scan() {
		text := scanner.Text()
		n := len(text)
		m := (n + 1) / 4
		if firstIteration {
			stacks = make([]Queue, m, m)
			firstIteration = false
			index := 1
			for i := 0; i < m; i++ {
				if (index < n) && strings.TrimSpace(string(text[index])) != "" {
					stacks[i].Append(string(text[index]))
				}
				index += 4
			}
			var numbersFound bool
			for !numbersFound {
				scanner.Scan()
				text = scanner.Text()
				if string(text[1]) == "1" {
					break
				}
				index = 1
				for i := 0; i < m; i++ {
					if (index < n) && strings.TrimSpace(string(text[index])) != "" {
						stacks[i].Append(string(text[index]))
					}
					index += 4
				}
			}
		} else {
			text = strings.TrimSpace(text)
			if text != "" {
				statement := strings.Split(text, " ")
				nMoves, err := strconv.Atoi(statement[1])
				if err != nil {
					log.Fatalln(err)
				}
				initialQueue, err := strconv.Atoi(statement[3])
				if err != nil {
					log.Fatalln(err)
				}
				finalQueue, err := strconv.Atoi(statement[5])
				if err != nil {
					log.Fatalln(err)
				}
				var cratesToTransport []string
				if nMoves > 1 {
					for i := 0; i < nMoves; i++ {
						crate := stacks[initialQueue-1].Pop()
						cratesToTransport = append(cratesToTransport, crate)
					}
					for i := nMoves - 1; i >= 0; i-- {
						crate := cratesToTransport[i]
						stacks[finalQueue-1].AppendLeft(crate)
					}
				} else {
					for i := 0; i < nMoves; i++ {
						crate := stacks[initialQueue-1].Pop()
						cratesToTransport = append(cratesToTransport, crate)
					}
					for i := 0; i < nMoves; i++ {
						crate := cratesToTransport[i]
						stacks[finalQueue-1].AppendLeft(crate)
					}
				}
			}
		}
	}
	var output string
	for _, v := range stacks {
		output += v.Peak()
	}
	fmt.Println(output)
}

func main() {
	partA()
	partB()

}
