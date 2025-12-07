package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	today "github.com/boris1592/aoc-2025/day2"
)

func main() {
	inputBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var (
		input = strings.TrimSpace(string(inputBytes))
		got1  = make(chan struct{})
		got2  = make(chan struct{})
	)

	go func() {
		defer func() { got1 <- struct{}{} }()

		ans, err := today.Solve1(input)
		if err != nil {
			fmt.Printf("part 1 got err: %s\n", err.Error())
		}

		fmt.Printf("part1 got: %v\n", ans)
	}()

	go func() {
		defer func() { got2 <- struct{}{} }()

		ans, err := today.Solve2(input)
		if err != nil {
			fmt.Printf("part 2 got err: %s", err.Error())
		}

		<-got1
		fmt.Printf("part2 got: %v\n", ans)
	}()

	<-got2
}
