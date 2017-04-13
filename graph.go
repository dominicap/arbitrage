package arbitrage

import (
	"encoding/json"
	"io/ioutil"
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
	Vertices  int
	Edges     int
	Adjacency []Edge
}

func (graph *Graph) addEdge(edge Edge) {
	graph.Adjacency = append(graph.Adjacency, edge)
	graph.Edges++
}

// BellmanFord represents a Bellman-Ford struct for use with the algorithm.
type BellmanFord struct {
	Graph     Graph
	Vertices  int
	Edges     int
	Distances []float64
	Cycle     []Edge
}

func (bellmanFord *BellmanFord) relax() {
	for i := 1; i <= bellmanFord.Vertices-1; i++ {
		for j := 0; j < bellmanFord.Edges; j++ {
			u := bellmanFord.Graph.Adjacency[j].Start
			v := bellmanFord.Graph.Adjacency[j].Destination

			weight := bellmanFord.Graph.Adjacency[j].Weight

			if bellmanFord.Distances[u]+weight < bellmanFord.Distances[v] {
				bellmanFord.Distances[v] = bellmanFord.Distances[u] + weight
			}

		}
	}
}

func (bellmanFord *BellmanFord) hasNegativeCycle() bool {
	flag := false
	for i := 0; i < bellmanFord.Edges; i++ {
		u := bellmanFord.Graph.Adjacency[i].Start
		v := bellmanFord.Graph.Adjacency[i].Destination
		weight := bellmanFord.Graph.Adjacency[i].Weight
		if bellmanFord.Distances[u]+weight < bellmanFord.Distances[v] {
			bellmanFord.Cycle = append(bellmanFord.Cycle, bellmanFord.Graph.Adjacency[i])
			flag = true
		}
	}
	return flag
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
