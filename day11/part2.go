package day11

import "strings"

func dfsExclusive(current, target, exclude string, graph map[string][]string, cache map[string]int) int {
	switch current {
	case exclude:
		return 0
	case target:
		return 1
	}

	if hit, ok := cache[current]; ok {
		return hit
	}

	var total int

	for _, next := range graph[current] {
		total += dfs(next, target, graph, cache)
	}

	cache[current] = total
	return total
}

func Solve2(input string) (any, error) {
	graph := map[string][]string{}

	for line := range strings.SplitSeq(input, "\n") {
		var (
			from = strings.Split(line, ":")[0]
			tos  = strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), " ")
		)

		for _, to := range tos {
			graph[to] = append(graph[to], from)
		}
	}

	var (
		svrToDac = dfsExclusive("dac", "svr", "fft", graph, map[string]int{})
		svrToFft = dfsExclusive("fft", "svr", "dac", graph, map[string]int{})
		dacToFft = dfsExclusive("fft", "dac", "", graph, map[string]int{})
		fftToDac = dfsExclusive("dac", "fft", "", graph, map[string]int{})
		fftToOut = dfsExclusive("out", "fft", "dac", graph, map[string]int{})
		dacToOut = dfsExclusive("out", "dac", "fft", graph, map[string]int{})

		ans = svrToDac*dacToFft*fftToOut + svrToFft*fftToDac*dacToOut
	)

	return ans, nil
}
