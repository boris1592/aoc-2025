package day5

import (
	"strconv"
	"strings"
)

func Solve1(input string) (any, error) {
	type idRange struct {
		from, to int
	}

	var (
		inputSplit = strings.Split(input, "\n\n")
		rangesStr  = strings.Split(inputSplit[0], "\n")
		idsStr     = strings.Split(inputSplit[1], "\n")

		ranges = make([]idRange, 0, len(rangesStr))
		anser  int
	)

	for _, rangeStr := range rangesStr {
		bounds := strings.Split(rangeStr, "-")
		from, err := strconv.Atoi(bounds[0])
		if err != nil {
			return nil, err
		}

		to, err := strconv.Atoi(bounds[1])
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, idRange{
			from: from,
			to:   to,
		})
	}

	for _, idStr := range idsStr {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}

		for _, r := range ranges {
			if id >= r.from && id <= r.to {
				anser++
				break
			}
		}
	}

	return anser, nil
}
