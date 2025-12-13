package day8

import (
	"slices"
	"strconv"
	"strings"
)

func findSize(start int, graph [][]int, used []bool) (total int) {
	if used[start] {
		return 0
	}

	used[start] = true
	total += 1

	for _, to := range graph[start] {
		total += findSize(to, graph, used)
	}

	return
}

// Please excuse my highly suboptimal solution, it was the first thing
// that came to mind.
func Solve1(input string) (any, error) {
	type box struct {
		x, y, z int
	}

	var (
		boxes []box
		graph [][]int
	)

	for line := range strings.SplitSeq(input, "\n") {
		xyz := strings.Split(line, ",")
		x, err := strconv.Atoi(xyz[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(xyz[1])
		if err != nil {
			return nil, err
		}

		z, err := strconv.Atoi(xyz[2])
		if err != nil {
			return nil, err
		}

		graph = append(graph, nil)
		boxes = append(boxes, box{x: x, y: y, z: z})
	}

	gotConns := make(map[int]map[int]struct{}, len(boxes))

	for range 1000 {
		var (
			minDist int
			first   = -1
			second  = -1
		)

		for i, a := range boxes {
			if _, ok := gotConns[i]; !ok {
				gotConns[i] = make(map[int]struct{})
			}

			for j, b := range boxes {
				if j <= i {
					continue
				}

				if _, ok := gotConns[i][j]; ok {
					continue
				}

				dist := distSquared(a.x, a.y, a.z, b.x, b.y, b.z)
				if first == -1 || dist < minDist {
					first = i
					second = j
					minDist = dist
				}
			}
		}

		if first == -1 {
			break
		}

		gotConns[first][second] = struct{}{}
		graph[first] = append(graph[first], second)
		graph[second] = append(graph[second], first)
	}

	var (
		sizes []int
		used  = make([]bool, len(boxes))
	)

	for i := range boxes {
		total := findSize(i, graph, used)
		if total > 0 {
			sizes = append(sizes, total)
		}
	}

	slices.Sort(sizes)

	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3], nil
}
