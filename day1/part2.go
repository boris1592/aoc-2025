package day1

import (
	"strconv"
	"strings"
)

func Solve2(input string) (any, error) {
	var (
		lines       = strings.Split(input, "\n")
		current int = 50
		total   int
	)

	for _, line := range lines {
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}

		if line[0] == 'L' {
			if current-rotation > 0 {
				current -= rotation
			} else {
				if current == 0 {
					total--
				}

				rotation -= current
				current = ((-rotation % 100) + 100) % 100
				total = total + 1 + rotation/100
			}
		} else {
			if current+rotation < 100 {
				current += rotation
			} else {
				rotation -= 100 - current
				current = rotation % 100
				total = total + 1 + rotation/100
			}
		}
	}

	return total, nil
}
