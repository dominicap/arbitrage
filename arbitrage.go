package arbitrage

import (
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

	for i := 0; i < total; i++ {
		for j := 0; j < total; j++ {
			rate := table[codes[i]][codes[j]]
			directedEdge := DirectedEdge{V: i, W: j, Weight: -math.Log(rate)}
		}
	}
	return nil
}
