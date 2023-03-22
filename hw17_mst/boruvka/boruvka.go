package boruvka

import (
	"fmt"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
	uf "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst/union-find"
)

func FindMST(g *graph.Graph) []graph.Edge {
	// visited := make(map[graph.Edge]bool)

	trees := uf.Init(g.GetVertices())
	fmt.Printf("Unions - vertex:root %+v\n", trees.GetPairs())

	for roots := trees.GetRoots(); len(roots) > 1; roots = trees.GetRoots() {
		fmt.Printf("Roots: %+v\n", roots)
		cheapest := make([]graph.Edge, len(roots))
		for i, root := range roots {
			cheapest[i] = g.CheapestEdge(root)
		}
		fmt.Printf("Cheapest edges: %+v\n", cheapest)

		for _, edge := range cheapest {
			trees.Union(edge.Src, edge.Dst)
		}
	}

	return nil
}
