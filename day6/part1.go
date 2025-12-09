package day6

import (
	"strconv"
	"strings"
)

func Solve1(input string) (any, error) {
	var (
		lines   = strings.Split(input, "\n")
		numbers = make([][]int, 0, len(lines)-1)
		ops     = make([]bool, 0)
	)

	for _, rawNumbers := range lines[:len(lines)-1] {
		var (
			numsSplit = strings.Fields(rawNumbers)
			row       = make([]int, 0, len(numbers))
		)

		for _, rawNumber := range numsSplit {
			num, err := strconv.Atoi(rawNumber)
			if err != nil {
				return nil, err
			}

			row = append(row, num)
		}

		numbers = append(numbers, row)
	}

	for op := range strings.FieldsSeq(lines[len(lines)-1]) {
		ops = append(ops, op[0] == '*')
	}

	var total int

	for i, op := range ops {
		var ans int

		if op {
			ans = 1
		}

		for _, row := range numbers {
			if op {
				ans *= row[i]
			} else {
				ans += row[i]
			}
		}

		total += ans
	}

	return total, nil
}
