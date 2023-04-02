package dijkstra

import (
	"math"

	bst "github.com/komarovn654/OTUS_Alg_Hw/bst"
	graph "github.com/komarovn654/OTUS_Alg_Hw/hw18_shortestpath"
)

type path struct {
	weight int
	edges  []graph.Edge
}

type pathways []path

func initPathways(len int) pathways {
	p := make([]path, len)

	for i := range p {
		p[i].weight = math.MaxInt
		p[i].edges = []graph.Edge{{Weight: math.MaxInt, Src: i, Dst: i}}
	}

	return p
}

func FindPath(g *graph.Graph, src int, dst int) []graph.Edge {
	// path := initPathways(len(g.GetAdjacentVectors()))

	// reachable := stack.New()
	// for _, vertex := range g.GetAdjacentVectors()[0] {
	// 	reachable.Push(vertex)
	// }

	// for vertex, ok := reachable.Pop(); ok; vertex, ok = reachable.Pop() {
	// 	v, ok := vertex.(graph.AdjacentVertex)
	// 	if !ok {
	// 		panic("type cast error")
	// 	}

	// 	for _, vertex := range g.GetAdjacentVectors()[v.Dst] {
	// 		reachable.Push(vertex)
	// 	}
	// 	fmt.Println(vertex)
	// }

	// return path[src].edges
	bst.InitBst()

	return nil
}

func add(x int, y int) int {
	if x > 0 && y > 0 && x+y < 0 {
		return math.MaxInt
	}
	return x + y
}
