package kruskal

import (
	"fmt"
	"sort"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
	uf "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst/union-find"
)

func FindMST(g *graph.Graph) (mst []graph.Edge) {
	vertices := g.GetVertices()
	mst = make([]graph.Edge, len(vertices)-1)

	// Get the edges and sort it by weights
	edges := sortEdges(g)

	// Save vertices in the union-find structure.
	pairs := uf.Init(vertices)

	// Save edges without cycles and with the lowest weights
	i := 0
	for _, edge := range edges {
		if pairs.Find(edge.Src) != pairs.Find(edge.Dst) {
			pairs.Union(edge.Src, edge.Dst)
			mst[i] = edge
			i++
			fmt.Println(mst)
		}
	}

	return mst
}

func sortEdges(g *graph.Graph) []graph.Edge {
	edges := g.GetEdges()
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	return edges
}
