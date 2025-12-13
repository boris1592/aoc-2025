package day9

import (
	"strconv"
	"strings"
)

func Solve1(input string) (any, error) {
	type point struct{ x, y int }

	points := make([]point, 0)

	for line := range strings.SplitSeq(input, "\n") {
		var (
			split = strings.Split(line, ",")
			point point
			err   error
		)

		point.x, err = strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}

		point.y, err = strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}

		points = append(points, point)
	}

	var largest int

	for i, p1 := range points {
		for _, p2 := range points[i+1:] {
			largest = max(
				largest,
				(1+max(p1.x, p2.x)-min(p1.x, p2.x))*(1+max(p1.y, p2.y)-min(p1.y, p2.y)),
			)
		}
	}

	return largest, nil
}
