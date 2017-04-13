package arbitrage

import (
	"math"
	"sort"
	"strings"
)

// Arbitrage returns a slice of ISO codes that yields the most profit.
func Arbitrage(value float64, code string) []string {
	codes, names := values()

	sort.Strings(codes)
	sort.Strings(names)

	base := code

	if value == 0 {
		panic("error: value is undefined or is 0.")
	}

	source := -1
	for index, code := range codes {
		if strings.EqualFold(code, base) {
			source = index
		}
	}

	if source == -1 {
		panic("error: ISO code not found.")
	}

	total := len(codes)

	table := createTable()

	graph := Graph{Vertices: total, Edges: 0}

	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			var rate float64
			if codes[i] == codes[j] {
				rate = 1
			} else {
				rate = table[codes[i]][codes[j]]
			}
			edge := Edge{Start: i, Destination: j, Weight: -math.Log(rate)}
			graph.addEdge(edge)
		}
	}

	bellmanFord := BellmanFord{Graph: graph, Distance: make([]float64, graph.Vertices), Predecessor: make([]int, graph.Vertices)}

	bellmanFord.initialize(source)
	bellmanFord.relax()

	if bellmanFord.hasNegativeCycle() {
		// TODO: Retrace and find shortest cycle with least weight
	}

	return nil
}
