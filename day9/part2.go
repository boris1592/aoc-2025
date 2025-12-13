package day9

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

func Solve2(input string) (any, error) {
	type point struct{ x, y int }

	var pointKey = func(p point) string {
		return fmt.Sprintf("%d,%d", p.x, p.y)
	}

	var (
		points  []point
		uniqueX = map[int]struct{}{}
		uniqueY = map[int]struct{}{}
	)

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

		uniqueX[point.x] = struct{}{}
		uniqueY[point.y] = struct{}{}
		points = append(points, point)
	}

	var (
		sortedX     = slices.Sorted(maps.Keys(uniqueX))
		sortedY     = slices.Sorted(maps.Keys(uniqueY))
		compressedX = make(map[int]int, len(sortedX))
		compressedY = make(map[int]int, len(sortedY))
	)

	for i, x := range sortedX {
		compressedX[x] = i + 1
	}

	for i, y := range sortedY {
		compressedY[y] = i + 1
	}

	border := map[string]struct{}{}

	for i, curr := range points {
		next := points[(i+1)%len(points)]

		if next.x < curr.x {
			next.x, curr.x = curr.x, next.x
		}

		if next.y < curr.y {
			next.y, curr.y = curr.y, next.y
		}

		if curr.x == next.x {
			for y := compressedY[curr.y]; y <= compressedY[next.y]; y++ {
				border[pointKey(point{x: compressedX[curr.x], y: y})] = struct{}{}
			}
		} else {
			for x := compressedX[curr.x]; x <= compressedX[next.x]; x++ {
				border[pointKey(point{x: x, y: compressedY[curr.y]})] = struct{}{}
			}
		}
	}

	var (
		outside = map[string]struct{}{pointKey(point{}): {}}
		queue   = []point{{}}
	)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dx := range []int{-1, 0, 1} {
			for _, dy := range []int{-1, 0, 1} {
				var (
					next         = point{x: curr.x + dx, y: curr.y + dy}
					nextKey      = pointKey(next)
					_, isBorder  = border[nextKey]
					_, isOutside = outside[nextKey]
				)

				if next.x < 0 || next.x > len(uniqueX)+1 ||
					next.y < 0 || next.y > len(uniqueY)+1 ||
					isBorder || isOutside {
					continue
				}

				outside[nextKey] = struct{}{}
				queue = append(queue, next)
			}
		}
	}

	var maxArea int

	for i, first := range points {
		for _, second := range points[i+1:] {
			var (
				first = first
				isBad bool
			)

			if second.x < first.x {
				second.x, first.x = first.x, second.x
			}

			if second.y < first.y {
				second.y, first.y = first.y, second.y
			}

		outerY:
			for _, y := range []int{compressedY[first.y], compressedY[second.y]} {
				for x := compressedX[first.x]; x <= compressedX[second.x]; x++ {
					key := pointKey(point{x: x, y: y})

					if _, ok := outside[key]; ok {
						isBad = true
						break outerY
					}
				}
			}

		outerX:
			for _, x := range []int{compressedX[first.x], compressedX[second.x]} {
				for y := compressedY[first.y]; y <= compressedY[second.y]; y++ {
					key := pointKey(point{x: x, y: y})

					if _, ok := outside[key]; ok {
						isBad = true
						break outerX
					}
				}
			}

			if isBad {
				continue
			}

			area := (second.y - first.y + 1) * (second.x - first.x + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea, nil
}
