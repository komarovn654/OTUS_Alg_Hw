package boruvka

import (
	"math"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
	uf "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst/union-find"
)

func initCheapestEdges(roots []int) map[int]graph.Edge {
	cheapest := make(map[int]graph.Edge) // tree's cheapest edge
	for _, root := range roots {
		cheapest[root] = graph.Edge{Weight: math.MaxInt}
	}

	return cheapest
}

func FindMST(g *graph.Graph) []graph.Edge {
	visited := make(map[graph.Edge]bool)

	trees := uf.Init(g.GetVertices())

	for roots := trees.GetRoots(); len(roots) > 1; roots = trees.GetRoots() {
		cheapest := initCheapestEdges(roots)

		// update cheapest edges map
		for vertex := range trees.GetPairs() {
			edge := g.CheapestEdge(vertex, visited)
			if edge.Weight < cheapest[trees.Find(vertex)].Weight {
				cheapest[trees.Find(vertex)] = edge
			}
		}

		// merging vertices in the cheapest edges
		for _, edge := range cheapest {
			trees.Union(edge.Src, edge.Dst)
			visited[edge] = true
		}
	}

	mst := make([]graph.Edge, 0)
	for edge := range visited {
		mst = append(mst, edge)
		deleteEdge(visited, edge, true)
	}
	return mst
}

func deleteEdge(edges map[graph.Edge]bool, key graph.Edge, deleteReverse bool) {
	delete(edges, key)
	if deleteReverse {
		reverse := graph.Edge{Weight: key.Weight, Src: key.Dst, Dst: key.Src}
		delete(edges, reverse)
	}
}
