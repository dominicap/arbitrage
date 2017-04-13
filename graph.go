package arbitrage

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
)

// Edge represents a edge from start to destination along with a weight.
type Edge struct {
	Start       int
	Destination int
	Weight      float64
}

// Graph represents a graph of vertices named 0 through V - 1.
type Graph struct {
	Vertices int
	Edges    int
	Edge     []Edge
}

func (graph *Graph) addEdge(edge Edge) {
	graph.Edge = append(graph.Edge, edge)
	graph.Edges++
}

// BellmanFord represents a Bellman-Ford struct for use with the algorithm.
type BellmanFord struct {
	Graph       Graph
	Distance    []float64
	Predecessor []int
}

func (bellmanFord *BellmanFord) initialize(source int) {
	for i := 0; i < bellmanFord.Graph.Vertices; i++ {
		bellmanFord.Distance[i] = math.Inf(+1)
		bellmanFord.Predecessor[i] = -1
	}
	bellmanFord.Distance[source] = 0
}

func (bellmanFord *BellmanFord) relax() {
	for i := 1; i <= bellmanFord.Graph.Vertices-1; i++ {
		for j := 0; j < bellmanFord.Graph.Edges; j++ {
			u := bellmanFord.Graph.Edge[j].Start
			v := bellmanFord.Graph.Edge[j].Destination
			weight := bellmanFord.Graph.Edge[j].Weight

			if bellmanFord.Distance[u]+weight < bellmanFord.Distance[v] {
				bellmanFord.Distance[v] = bellmanFord.Distance[u] + weight
				bellmanFord.Predecessor[v] = u
			}
		}
	}
}

func (bellmanFord *BellmanFord) hasNegativeCycle() bool {
	for j := 0; j < bellmanFord.Graph.Edges; j++ {
		u := bellmanFord.Graph.Edge[j].Start
		v := bellmanFord.Graph.Edge[j].Destination
		weight := bellmanFord.Graph.Edge[j].Weight

		if bellmanFord.Distance[u]+weight < bellmanFord.Distance[v] {
			return true
		}
	}
	return false
}

func (bellmanFord *BellmanFord) retraceNegativeCycle(predecessor []Edge, source int) []Edge {
	return make([]Edge, 1) // Stub
}

func createTable() map[string]map[string]float64 {
	codes, _ := values()
	sort.Strings(codes)

	table := make(map[string]map[string]float64)

	for _, code := range codes {
		url := latestURL + "?base=" + code

		response, err := http.Get(url)
		check(err)
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		check(err)

		var data LatestExchangeData
		json.Unmarshal(body, &data)

		table[data.Base] = data.Rates
	}

	return table
}
