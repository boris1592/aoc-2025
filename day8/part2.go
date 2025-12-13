package day8

import (
	"cmp"
	"slices"
	"strconv"
	"strings"
)

type dsu struct {
	parent map[int]int
}

func newDsu() dsu {
	return dsu{parent: map[int]int{}}
}

func (d *dsu) find(a int) int {
	if p, ok := d.parent[a]; ok {
		p = d.find(p)
		d.parent[a] = p
		return p
	}

	return a
}

func (d *dsu) check(a, b int) bool {
	return d.find(a) == d.find(b)
}

func (d *dsu) add(a, b int) {
	if d.check(a, b) {
		return
	}

	d.parent[d.find(a)] = d.find(b)
}

func Solve2(input string) (any, error) {
	type box struct {
		x, y, z int
	}

	type pair struct {
		a, b int
	}

	var (
		lines = strings.Split(input, "\n")
		boxes = make([]box, 0, len(lines))
		pairs = make([]pair, 0, len(lines)*(len(lines)+1)/2)
	)

	for _, line := range lines {
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

		boxes = append(boxes, box{x: x, y: y, z: z})
	}

	for i := range len(boxes) {
		for j := range len(boxes) - i {
			j += i
			pairs = append(pairs, pair{a: i, b: j})
		}
	}

	slices.SortFunc(pairs, func(a, b pair) int {
		var (
			distA = distSquared(boxes[a.a].x, boxes[a.a].y, boxes[a.a].z, boxes[a.b].x, boxes[a.b].y, boxes[a.b].z)
			distB = distSquared(boxes[b.a].x, boxes[b.a].y, boxes[b.a].z, boxes[b.b].x, boxes[b.b].y, boxes[b.b].z)
		)

		return cmp.Compare(distA, distB)
	})

	dsu := newDsu()

	for _, pair := range pairs {
		dsu.add(pair.a, pair.b)

		done := true

		for i := range len(boxes) {
			if !dsu.check(0, i) {
				done = false
				break
			}
		}

		if done {
			return boxes[pair.a].x * boxes[pair.b].x, nil
		}
	}

	return nil, nil
}
