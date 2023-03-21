package boruvka

import (
	"fmt"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
	uf "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst/union-find"
)

func FindMST(g *graph.Graph) []graph.Edge {
	roots := make(map[any]graph.Edge) // {union's root: min weight union's edge}
	set := uf.Init(g.GetVertices())

	set.Union(0, 1)
	set.Union(2, 3)

	// av := g.GetAdjacentVectors()
	for vertex := range set.GetPairs() {
		fmt.Printf("Vertex: %v, Vertex min edge: %+v, Vertex root: %v\n", vertex, g.MinWeightEdge(vertex.(int)), set.Find(vertex))
		if _, ok := roots[set.Find(vertex)]; !ok {
			roots[set.Find(vertex)] = g.MinWeightEdge(vertex.(int))
			continue
		}
		if g.MinWeightEdge(vertex.(int)).Weight < roots[set.Find(vertex)].Weight {
			roots[set.Find(vertex)] = g.MinWeightEdge(vertex.(int))
		}
	}

	for key := range roots {
		fmt.Printf("Root: %v, Edge: %+v\n", key, roots[key])
	}

	return nil
}
