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
		fmt.Println(base)
		panic("error: ISO code not found.")
	}

	total := len(codes)

	table := createTable()

	graph := new(EdgeWeightedDigraph)

	graph.V = total
	graph.E = 0
	graph.InDegree = make([]int, total)

	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			var rate float64
			if codes[i] == codes[j] {
				rate = 1
			} else {
				rate = table[codes[i]][codes[j]]
			}
			directedEdge := DirectedEdge{V: i, W: j, Weight: -math.Log(rate)}
			graph.Adjacency = append(graph.Adjacency, directedEdge)
		}
	}

	bellmanFord := new(BellmanFord)

	bellmanFord.DistanceTo = make([]float64, graph.V)
	bellmanFord.EdgeTo = make([]DirectedEdge, graph.V)
	bellmanFord.OnQueue = make([]bool, graph.V)

	for i := 0; i < graph.V; i++ {
		bellmanFord.DistanceTo[i] = math.Inf(+1)
	}

	bellmanFord.DistanceTo[source] = 0.0

	return nil
}
