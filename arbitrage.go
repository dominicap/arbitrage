package arbitrage

import (
	"fmt"
	"math"
	"sort"
)

// Arbitrage returns a slice of ISO codes that yields the most profit.
func Arbitrage() []string {
	codes, names := values()

	sort.Strings(codes)
	sort.Strings(names)

	total := len(codes)

	table := createTable()

	graph := new(EdgeWeightedDigraph)

	graph.V = total
	graph.E = 0

	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			if !(codes[i] == codes[j]) {
				rate := table[codes[i]][codes[j]]
				directedEdge := DirectedEdge {V: i, W: j, Weight: -math.Log(rate)}
				graph.Adjacency = append(graph.Adjacency, directedEdge)
			}
		}
	}

	return nil
}
