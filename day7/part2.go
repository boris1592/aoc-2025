package day7

import "strings"

func Solve2(input string) (any, error) {
	type cellKind int

	type cell struct {
		kind  cellKind
		paths int
	}

	const (
		empty cellKind = iota + 1
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
				cell.kind = empty
			case 'S':
				cell.kind = start
			case '^':
				cell.kind = splitter
			}

			row = append(row, cell)
		}

		field = append(field, row)
	}

	for row := range field {
		for col := range len(field[row]) {
			switch field[row][col].kind {
			case start:
				field[row+1][col].kind = beam
				field[row+1][col].paths++
			case beam:
				if row >= len(field)-1 {
					continue
				}

				// dynamic programming 😳
				switch field[row+1][col].kind {
				case splitter:
					field[row+1][col-1].kind = beam
					field[row+1][col-1].paths += field[row][col].paths

					field[row+1][col+1].kind = beam
					field[row+1][col+1].paths += field[row][col].paths

				default:
					field[row+1][col].kind = beam
					field[row+1][col].paths += field[row][col].paths
				}
			}
		}
	}

	for _, cell := range field[len(field)-1] {
		total += cell.paths
	}

	return total, nil
}
