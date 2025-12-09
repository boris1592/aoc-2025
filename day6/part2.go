package day6

import "strings"

// I hate this. It's just an exercise in careful parsing.
func Solve2(input string) (any, error) {
	var (
		lines = strings.Split(input, "\n")
		ops   = lines[len(lines)-1]
		curr  int
		sum   int
	)

	for curr < len(ops) {
		var (
			op    = ops[curr]
			till  = curr + 1
			total int
		)

		for till < len(ops) && ops[till] == ' ' {
			till++
		}

		if till == len(ops) {
			for _, l := range lines {
				till = max(till, len(l)+1)
			}
		}

		if op == '*' {
			total = 1
		}

		for ; curr < till-1; curr++ {
			var num int

			for row := 0; row < len(lines)-1; row++ {
				if curr >= len(lines[row]) || lines[row][curr] == ' ' {
					continue
				}

				num *= 10
				num += int(lines[row][curr] - '0')
			}

			if op == '*' {
				total *= num
			} else {
				total += num
			}
		}

		sum += total
		curr++
	}

	return sum, nil
}
