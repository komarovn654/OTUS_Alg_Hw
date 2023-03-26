package prim

import (
	"fmt"
	"math"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw17_mst"
)

type turnStat struct {
	src    int // source Vertex number
	weight int // turn's weight from source Vertex
}

type nextTurn map[int]turnStat // [destination vertex]: {source vertex, weight}

// cheapest edges
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

// func FindMST(g *graph.Graph) []graph.Edge {
// 	mst := make([]graph.Edge, 0)
// 	vertices := g.GetVertices()
// 	visited := make(map[int]bool) // visited vertices
// 	visited[0] = true

// 	cheapest := make(map[int]graph.Edge) // tree's cheapest edge
// 	for _, vertex := range vertices {
// 		cheapest[vertex] = graph.Edge{Weight: math.MaxInt}
// 	}

// 	for srcVertNum, vector := range g.GetAdjacentVectors() {
// 		for _, dstVert := range vector {
// 			edge := graph.Edge{Weight: dstVert.Weight, Src: srcVertNum, Dst: dstVert.Num}
// 			if !visited[dstVert.Num] && (dstVert.Weight < turns[dstVert.Num].weight) {
// 				turns[dstVert.Num] = turnStat{weight: dstVert.Weight, src: srcVert}
// 			}
// 		}

// 		edge := turns.getMinWeightEdge()
// 		visited[edge.Dst] = true
// 		delete(turns, edge.Dst)

// 		mst = append(mst, edge)
// 		if len(mst) >= (len(vertices) - 1) {
// 			break
// 		}
// 	}

// 	return mst
// }

type edgeStatus struct {
	edge    graph.Edge
	visited bool // src vertex visited status
}

type minSteps map[int]edgeStatus // destination vertex: edge status

func initSteps(vertices []int) minSteps {
	steps := make(map[int]edgeStatus)

	for _, vertex := range vertices {
		steps[vertex] = edgeStatus{edge: graph.Edge{Weight: math.MaxInt, Src: -1, Dst: vertex}}
	}

	steps[0] = edgeStatus{edge: graph.Edge{Weight: math.MaxInt, Src: 0, Dst: 0}, visited: true} // start vertex
	return steps
}

func (this *minSteps) findCheapest() graph.Edge {
	min := graph.Edge{Weight: math.MaxInt, Src: -1, Dst: -1}

	for _, step := range *this {
		if step.edge.Weight < min.Weight && !step.visited {
			min = step.edge
		}
	}

	return min
}

func (this *minSteps) updateWeights(g *graph.Graph) {
	for _, step := range *this {
		if step.visited {
			for _, adjVertex := range g.GetAdjacentVectors()[step.edge.Dst] {
				if adjVertex.Weight < (*this)[adjVertex.Num].edge.Weight {
					(*this)[adjVertex.Num] = edgeStatus{
						edge: graph.Edge{
							Weight: adjVertex.Weight,
							Src:    step.edge.Dst,
							Dst:    adjVertex.Num,
						},
						visited: (*this)[adjVertex.Num].visited,
					}
				}
			}
		}
	}
}

func FindMST(g *graph.Graph) []graph.Edge {
	mst := make([]graph.Edge, 0)

	steps := initSteps(g.GetVertices())

	for i := 0; i < len(g.GetAdjacentVectors())-1; i++ {
		steps.updateWeights(g)

		edge := steps.findCheapest()

		steps[edge.Dst] = edgeStatus{
			edge:    graph.Edge{Weight: edge.Weight, Src: edge.Src, Dst: edge.Dst},
			visited: true,
		}

		mst = append(mst, edge)
		fmt.Println(mst)
	}

	return mst
}
