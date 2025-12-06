package day1

import (
	"strconv"
	"strings"
)

func Solve1(input string) (any, error) {
	var (
		lines   = strings.Split(input, "\n")
		current = 50
		total   int
	)

	for _, line := range lines {
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		if line[0] == 'L' {
			rotation = -rotation
		}

		current += rotation
		current = ((current % 100) + 100) % 100

		if current == 0 {
			total++
		}
	}

	return total, nil
}
