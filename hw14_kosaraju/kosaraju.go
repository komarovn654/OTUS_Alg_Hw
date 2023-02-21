package hw14_kosaraju

import "fmt"

const (
	FIRST_VERTEX = 0
)

type Vertices []int
type Vertex int
type Graph []Vertices // Vertices nums begins with 0

type kosarajuGraph struct {
	graph    Graph
	vertices stack
	visited  tarjanVerts // { vertex: tarjan's color }
}

func (kg *kosarajuGraph) Kosaraju() {
	invert := kosarajuGraph{graph: *kg.graph.reverse()}

	stack := invert.TarjanDFS(FIRST_VERTEX)
	stack.forEach(func(a any) {
		fmt.Println(a)
	})
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

	// reverse graph
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
func (kg *kosarajuGraph) TarjanDFS(start Vertex) stack {
	if kg.graph == nil {
		return stack{}
	}

	return kg.DFS(start)
}
