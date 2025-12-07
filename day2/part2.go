package day2

import (
	"math"
	"strconv"
	"strings"
)

func solveForRange2(from, to int) (total int) {
	used := make(map[int]struct{})

	for divisor := 2; divisor <= intLen(to); divisor++ {
		currLen := intLen(from)
		currLen += (divisor - currLen%divisor) % divisor

		for currLen <= intLen(to) {
			var step int
			for i := range divisor {
				step += int(math.Pow10(i * currLen / divisor))
			}

			var (
				lower = (int(math.Pow10(currLen-1))/step + 1) * step
				upper = (int(math.Pow10(currLen)) / step) * step
			)

			if lower < from {
				lower = (from/step + 1) * step
			}

			if upper > to {
				upper = (to / step) * step
			}

			for i := lower; i <= upper; i += step {
				if _, ok := used[i]; ok {
					continue
				}

				total += i
				used[i] = struct{}{}
			}

			currLen += divisor
		}
	}

	return
}

func Solve2(input string) (any, error) {
	var total int

	for numRange := range strings.SplitSeq(input, ",") {
		nums := strings.Split(numRange, "-")

		from, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}

		to, err := strconv.Atoi(nums[1])
		if err != nil {
			return nil, err
		}

		total += solveForRange2(from, to)
	}

	return total, nil
}
