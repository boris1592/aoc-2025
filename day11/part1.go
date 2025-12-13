package day11

import "strings"

func dfs(current, target string, graph map[string][]string, cache map[string]int) int {
	if current == target {
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

func Solve1(input string) (any, error) {
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

	return dfs("out", "you", graph, map[string]int{}), nil
}
