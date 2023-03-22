package prim

import (
	"math"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
)

type turnStat struct {
	src    int // source Vertex number
	weight int // turn's weight from source Vertex
}

type nextTurn map[int]turnStat // [destination vertex]: {source vertex, weight}

// Init turns with max weigths
func initTurns(vertices []int) nextTurn {
	turns := make(nextTurn)
	for _, vertex := range vertices {
		turns[vertex] = turnStat{-1, math.MaxInt}
	}

	return turns
}

func (turn *nextTurn) getMinWeightEdge() (edge graph.Edge) {
	min := math.MaxInt
	for v := range *turn {
		if (*turn)[v].weight < min {
			min = (*turn)[v].weight
			edge.Weight = min
			edge.Src = (*turn)[v].src
			edge.Dst = v
		}
	}

	return
}

func FindMST(g *graph.Graph) []graph.Edge {
	mst := make([]graph.Edge, 0)
	vertices := g.GetVertices()
	visited := make(map[int]bool)

	turns := initTurns(vertices)
	visited[0] = true

	for vertex1, vector := range g.GetAdjacentVectors() {
		for _, vertex2 := range vector {
			if !visited[vertex2.Num] && (vertex2.Weight < turns[vertex2.Num].weight) {
				turns[vertex2.Num] = turnStat{weight: vertex2.Weight, src: vertex1}
			}
		}

		edge := turns.getMinWeightEdge()
		visited[edge.Dst] = true
		delete(turns, edge.Dst)

		mst = append(mst, edge)
		if len(mst) >= (len(vertices) - 1) {
			break
		}
	}

	return mst
}
