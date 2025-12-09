package day7

import "strings"

func Solve1(input string) (any, error) {
	type cell int

	const (
		empty cell = iota + 1
		start
		beam
		splitter
	)

	var (
		lines = strings.Split(input, "\n")
		field = make([][]cell, 0, len(lines))
		total int
	)

	for _, line := range lines {
		row := make([]cell, 0, len(line))

		for _, char := range line {
			var cell cell

			switch char {
			case '.':
				cell = empty
			case 'S':
				cell = start
			case '^':
				cell = splitter
			}

			row = append(row, cell)
		}

		field = append(field, row)
	}

	for row := range field {
		for col := range len(field[row]) {
			switch field[row][col] {
			case start:
				field[row+1][col] = beam
			case beam:
				if row >= len(field)-1 {
					continue
				}

				switch field[row+1][col] {
				case empty:
					field[row+1][col] = beam
				case splitter:
					field[row+1][col-1] = beam
					field[row+1][col+1] = beam
					total++
				}
			}
		}
	}

	return total, nil
}
