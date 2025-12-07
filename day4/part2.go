package day4

import (
	"strings"
)

func Solve2(input string) (any, error) {
	grid := make([][]bool, 0)

	for row := range strings.SplitSeq(input, "\n") {
		gridRow := make([]bool, 0, len(row))

		for _, c := range row {
			gridRow = append(gridRow, c == '@')
		}

		grid = append(grid, gridRow)
	}

	var (
		answer int
		past   = 1
	)

	for past > 0 {
		past = 0

		for row := range grid {
			for col := range grid[row] {
				if !grid[row][col] {
					continue
				}

				var total int

				for _, drow := range []int{-1, 0, 1} {
					for _, dcol := range []int{-1, 0, 1} {
						if drow == 0 && dcol == 0 {
							continue
						}

						var (
							r = row + drow
							c = col + dcol
						)

						if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) {
							continue
						}

						if grid[r][c] {
							total++
						}
					}
				}

				if total < 4 {
					answer++
					past++
					grid[row][col] = false
				}
			}
		}
	}

	return answer, nil
}
