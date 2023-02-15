package hw14_kosaraju

import "fmt"

const (
	FIRST_VERTEX = 0
)

type Vertices []int
type Vertex int
type Graph []Vertices // Vertices nums begins with 0

type ms struct {
	vertice int
	color   int
}

func (g *Graph) search() {
	invert := g.reverse()
	fmt.Println(invert.TarjanDFS(2))
}

// reverse graph
func (g *Graph) reverse() *Graph {
	if len(*g) == 0 {
		return g
	}

	// init empty graph
	invert := make(Graph, len(*g))
	for i := range invert {
		invert[i] = make(Vertices, 0)
	}

	// invert graph
	for i, vertice := range *g {
		for _, value := range vertice {
			if value == -1 {
				continue
			}
			invert[value] = append(invert[value], i)
		}
	}

	return &invert
}

// Tarjan depth-first search. Return map{visited vertex: tarjan's color}.
func (g *Graph) TarjanDFS(start Vertex) vertices {
	if g == nil {
		return vertices{}
	}

	tg := tarjanGraph{g: *g}
	return tg.DFS(start)
}
