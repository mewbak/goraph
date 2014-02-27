// Package spfw finds the all-pairs shortest paths
// using Floyd-Warshall algorithm.
package spfw

/*
http://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm

let dist be a |V| × |V| array of minimum distances initialized to ∞ (infinity)

for each vertex v
	dist[v][v] ← 0

for each edge (u,v)
	dist[u][v] ← w(u,v)  // the weight of the edge (u,v)

for k from 1 to |V|
	for i from 1 to |V|
		for j from 1 to |V|
			if dist[i][j] > dist[i][k] + dist[k][j]
				dist[i][j] ← dist[i][k] + dist[k][j]
			end if
*/

import "github.com/gyuho/goraph/graph/gt"

func SPFW(g *gt.Graph, src, dst string) (float64, string) {
	make2DSlice := func(row, column int) [][]float64 {
		mat := make([][]float64, row)
		for i := range mat {
			mat[i] = make([]float64, column)
		}
		return mat
	}
	// let dist be a |V| × |V| array of minimum distances initialized to ∞ (infinity)
	dist := make2DSlice(*g.Size, *g.Size)

	// to initialize
	for r := range dist {
		for c := range dist[r] {
			dist[r][c] = 9999999999.9999
		}
	}

	// for each vertex v
	// dist[v][v] ← 0
	vertices := g.GetVertices()
	for _, v := range vertices {
		dist[g.Index[v]][g.Index[v]] = 0
	}

	// for each edge (u,v)
	// if the edge weight is not 0.0
	// that means there is an edge
	for r := range g.Matrix {
		for c := range g.Matrix[r] {
			if g.Matrix[r][c] != 0.0 {
				// dist[u][v] ← w(u,v)
				dist[r][c] = g.Matrix[r][c]
			}
		}
	}

	/*
	   for k from 1 to |V|
	   	for i from 1 to |V|
	   		for j from 1 to |V|
	   			if dist[i][j] > dist[i][k] + dist[k][j]
	   				dist[i][j] ← dist[i][k] + dist[k][j]
	   			end if
	*/
	for k := 0; k < len(vertices); k++ {
		for i := 0; i < len(vertices); i++ {
			for j := 0; j < len(vertices); j++ {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	return dist[g.Index[src]][g.Index[dst]], gt.Output2DSlice(dist)
}
