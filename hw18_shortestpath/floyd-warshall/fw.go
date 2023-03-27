package floydwarshall

import (
	"fmt"
	"math"

	graph "github.com/komarovn654/OTUS_Alg_Hw/hw18_shortestpath"
)

type path struct {
	weight int
	edges  []graph.Edge
}

type pathMatrix [][]path

func initPathMatrix(len int) pathMatrix {
	matrix := make(pathMatrix, len)
	for i := range matrix {
		matrix[i] = make([]path, len)
		for j := range matrix[i] {
			if i == j {
				matrix[i][j].weight = 0
				matrix[i][j].edges = []graph.Edge{{Weight: 0, Src: i, Dst: j}}
				continue
			}
			matrix[i][j].weight = math.MaxInt
		}
	}

	return matrix
}

func (this *pathMatrix) setDirectPathes(g *graph.Graph) {
	for _, edge := range g.GetEdges() {
		(*this)[edge.Src][edge.Dst].weight = edge.Weight
		(*this)[edge.Src][edge.Dst].edges = []graph.Edge{edge}
	}
}

func (this *pathMatrix) print() {
	for src, line := range *this {
		for dst, path := range line {
			fmt.Printf("From %v to %v; Weigth: %v; Path: %+v\n", src, dst, path.weight, path.edges)
		}
		fmt.Println()
	}
}

func FindPath(g *graph.Graph, src int, dst int) []graph.Edge {
	matrix := initPathMatrix(len(g.GetAdjacentVectors()))
	vertices := g.GetVertices()

	matrix.setDirectPathes(g)

	for _, k := range vertices {
		for _, src := range vertices {
			for _, dst := range vertices {
				if add(matrix[src][k].weight, matrix[k][dst].weight) < matrix[src][dst].weight {
					matrix[src][dst].weight = add(matrix[src][k].weight, matrix[k][dst].weight)
					matrix[src][dst].edges = matrix[src][k].edges
					matrix[src][dst].edges = append(matrix[src][dst].edges, matrix[k][dst].edges...)
				}
			}
		}
	}
	// matrix.print()
	return matrix[src][dst].edges
}

func add(x int, y int) int {
	if x > 0 && y > 0 && x+y < 0 {
		return math.MaxInt
	}
	return x + y
}
