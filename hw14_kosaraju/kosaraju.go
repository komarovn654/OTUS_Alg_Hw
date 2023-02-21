package hw14_kosaraju

type Vertex int
type Vertices []Vertex
type Graph []Vertices // Vertices nums begins with 0

type kosarajuGraph struct {
	graph    Graph
	vertices stack
	visited  tarjanVerts // { vertex: tarjan's color }
}

// Return strongly connected components in ordered graph
func (g *Graph) Kosaraju() (strConnected []stack) {
	// init direct and reverse graphs
	kg := kosarajuGraph{graph: *g, visited: make(tarjanVerts), vertices: Stack()}
	invert := kosarajuGraph{graph: *g.reverse(), visited: make(tarjanVerts), vertices: Stack()}
	var order stack // vertices after DFS in reverse graph

	// get vertices via tarjan DFS in reverse graph
	for vertex := range invert.graph {
		if color, ok := invert.visited[Vertex(vertex)]; !ok || color != BLACK {
			order = invert.TarjanDFS(Vertex(vertex))
		}
	}

	// DFS in direct graph
	for {
		vertex, ok := order.Pop()
		if !ok {
			return strConnected // empty stack
		}

		v, ok := vertex.(Vertex)
		if !ok {
			panic("type cast error")
		}

		if color, ok := kg.visited[v]; !ok || color != BLACK {
			strConnected = append(strConnected, kg.TarjanDFS(v))
			kg.vertices = Stack() // refresh stack
		}
	}
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
			invert[value] = append(invert[value], Vertex(i))
		}
	}

	return &invert
}
