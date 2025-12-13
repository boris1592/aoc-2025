package day10

import (
	"strconv"
	"strings"
)

func bfs1(light uint16, switches []uint16) int {
	var (
		// This is absolutely suboptimal but I'm not writing a proper queue so here we go
		queue = []uint16{0}
		dist  = map[uint16]int{0: 0}
	)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, sw := range switches {
			next := curr ^ sw

			if _, ok := dist[next]; ok {
				continue
			}

			dist[next] = dist[curr] + 1
			queue = append(queue, next)
		}
	}

	return dist[light]
}

func Solve1(input string) (any, error) {
	var (
		lines = strings.Split(input, "\n")
		ans   int
	)

	for _, line := range lines {
		var (
			pieces   = strings.Split(line, " ")
			light    uint16
			switches = make([]uint16, len(pieces)-2)
		)

		for i, char := range pieces[0][1 : len(pieces[0])-1] {
			if char == '#' {
				light |= 1 << i
			}
		}

		for _, piece := range pieces[1 : len(pieces)-1] {
			var (
				bits = strings.Split(piece[1:len(piece)-1], ",")
				sw   uint16
			)

			for _, bitStr := range bits {
				bit, err := strconv.Atoi(bitStr)
				if err != nil {
					return nil, err
				}

				sw |= 1 << bit
			}

			switches = append(switches, sw)
		}

		steps := bfs1(light, switches)
		ans += steps
	}

	return ans, nil
}
