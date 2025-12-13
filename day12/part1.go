package day12

import (
	"strconv"
	"strings"
)

func Solve1(input string) (any, error) {
	// LMAO, I'm pretty used to AoC problems being almost unsolvable (or
	// literally unsolvable) in general case and the input containing some
	// little property that actually makes them solvable. But this is
	// something else. There's literally no way to find out about it
	// here without actually submitting an answer that you can't prove
	// is right.

	var (
		splitInput = strings.Split(input, "\n\n")
		areas      [6]int
		ans        int
	)

	for i, block := range splitInput[:6] {
		areas[i] = strings.Count(block, "#")
	}

	for region := range strings.SplitSeq(splitInput[6], "\n") {
		var (
			dimensions           = strings.Split(strings.Split(region, ":")[0], "x")
			quantities           = strings.Split(strings.TrimSpace(strings.Split(region, ":")[1]), " ")
			width, height, total int
			err                  error
		)

		width, err = strconv.Atoi(dimensions[0])
		if err != nil {
			return nil, err
		}

		height, err = strconv.Atoi(dimensions[1])
		if err != nil {
			return nil, err
		}

		for i, quantityRaw := range quantities {
			quantity, err := strconv.Atoi(quantityRaw)
			if err != nil {
				return nil, err
			}

			total += areas[i] * quantity
		}

		if total < width*height {
			ans++
		}
	}

	return ans, nil
}
