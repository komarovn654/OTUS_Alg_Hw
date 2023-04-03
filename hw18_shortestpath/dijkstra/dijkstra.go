package dijkstra

import (
	"math"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw18_shortestpath"
	"github.com/komarovn654/OTUS_Alg_Hw/hw18_shortestpath/queue"
)

type path struct {
	weight  int
	edges   []graph.Edge
	visited bool
}

type pathways []path

func initPathways(len int) pathways {
	p := make([]path, len)

	for i := range p {
		p[i].weight = math.MaxInt
	}

	return p
}

func FindPath(g *graph.Graph, src int, dst int) []graph.Edge {
	if src == dst {
		return []graph.Edge{{Src: src, Dst: dst, Weight: 0}}
	}
	shortest := initPathways(len(g.GetAdjacentVectors()))

	// Init the queue with reachable vertices from first vertex
	reachable := queue.NewQueue()
	for _, vertex := range g.GetAdjacentVectors()[src] {
		shortest[vertex.Dst].weight = vertex.Weight
		edge := graph.Edge{Src: src, Dst: vertex.Dst, Weight: vertex.Weight}
		shortest[vertex.Dst].edges = append(shortest[vertex.Dst].edges, edge)

		reachable.Queue(edge)
	}
	shortest[src].visited = true

	// while queue is not empty
	for edge, ok := reachable.Dequeue(); ok; edge, ok = reachable.Dequeue() {
		e, ok := edge.(graph.Edge)
		if !ok {
			panic("type cast error")
		}

		if add(shortest[e.Src].weight, e.Weight) < shortest[e.Dst].weight {
			shortest[e.Dst].weight = add(shortest[e.Src].weight, e.Weight)
			shortest[e.Dst].edges = shortest[e.Src].edges
			shortest[e.Dst].edges = append(shortest[e.Dst].edges, e)
		}

		for _, vertex := range g.GetAdjacentVectors()[e.Dst] {
			if !shortest[vertex.Dst].visited {
				reachable.Queue(graph.Edge{Src: e.Dst, Dst: vertex.Dst, Weight: vertex.Weight})
			}
		}
		shortest[e.Dst].visited = true
	}

	return shortest[dst].edges
}

func add(x int, y int) int {
	if x > 0 && y > 0 && x+y < 0 {
		return math.MaxInt
	}
	return x + y
}
