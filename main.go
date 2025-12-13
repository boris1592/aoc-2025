package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	today "github.com/boris1592/aoc-2025/day11"
)

func runSolution[F ~func(string) (any, error)](name string, solution F, input string) {
	var (
		start    = time.Now()
		ans, err = solution(input)
		duration = time.Since(start).Milliseconds()
	)

	if err != nil {
		fmt.Printf("%s got err: %s\n", name, err.Error())
	}

	fmt.Printf("%s got: %v in %d ms\n", name, ans, duration)
}

func main() {
	inputBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var (
		input = strings.TrimSpace(string(inputBytes))
		wg    sync.WaitGroup
	)

	wg.Go(func() {
		runSolution("part 1", today.Solve1, input)
	})

	wg.Go(func() {
		runSolution("part 2", today.Solve2, input)
	})

	wg.Wait()
}
