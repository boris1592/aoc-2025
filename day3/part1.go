package day3

import "strings"

func Solve1(input string) (any, error) {
	var total int

	for bank := range strings.SplitSeq(input, "\n") {
		var mx int

		for i, batt1 := range bank {
			for _, batt2 := range bank[i+1:] {
				sum := int(byte(batt1)-'0')*10 + int(byte(batt2)-'0')
				if sum > mx {
					mx = sum
				}
			}
		}

		total += mx
	}

	return total, nil
}
