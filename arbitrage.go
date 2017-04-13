package arbitrage

import (
	"fmt"
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

	bellmanFord := BellmanFord{Graph: graph, Vertices: graph.Vertices, Edges: graph.Edges, Distances: make([]float64, graph.Vertices)}

	for i := 0; i < bellmanFord.Vertices; i++ {
		bellmanFord.Distances[i] = math.Inf(+1)
	}

	bellmanFord.Distances[source] = 0

	bellmanFord.relax()

	if bellmanFord.hasNegativeCycle() {
		for i := 0; i < len(bellmanFord.Cycle); i++ {
			rate := table[codes[bellmanFord.Cycle[i].Start]][codes[bellmanFord.Cycle[i].Destination]]
			fmt.Println(codes[bellmanFord.Cycle[i].Start], codes[bellmanFord.Cycle[i].Destination], rate)
		}
	}

	return nil
}
