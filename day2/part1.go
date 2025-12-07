package day2

import (
	"math"
	"strconv"
	"strings"
)

func intLen(n int) (l int) {
	for n > 0 {
		l++
		n /= 10
	}

	return
}

func solveForRange(from, to int) (total int) {
	currLen := intLen(from)
	currLen += currLen % 2

	for currLen <= intLen(to) {
		var (
			step  = int(math.Pow10(currLen/2)) + 1
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
			total += i
		}

		currLen += 2
	}

	return
}

func Solve1(input string) (any, error) {
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

		total += solveForRange(from, to)
	}

	return total, nil
}
