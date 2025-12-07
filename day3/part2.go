package day3

import "strings"

func solveForSubstr(s string, l int, acc int) int {
	if l <= 0 {
		return acc
	}

	var (
		mx    int
		index int
	)

	for i, c := range s[:len(s)-l+1] {
		digit := int(byte(c) - '0')

		if digit > mx {
			mx = digit
			index = i
		}
	}

	return solveForSubstr(s[index+1:], l-1, acc*10+mx)
}

func Solve2(input string) (any, error) {
	var total int

	for line := range strings.SplitSeq(input, "\n") {
		total += solveForSubstr(line, 12, 0)
	}

	return total, nil
}
