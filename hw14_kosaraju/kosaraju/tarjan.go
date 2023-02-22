package hw14_kosaraju

const (
	WHITE = iota
	GREY
	BLACK
)

type tarjanVerts map[Vertex]int // { vertex: tarjan's color}

func (kg *kosarajuGraph) TarjanDFS(start Vertex) stack {
	if kg == nil {
		return kg.vertices
	}

	if len(kg.graph) == 0 {
		return kg.vertices
	}

	kg.tarjanDFS(start)
	return kg.vertices
}

func (kg *kosarajuGraph) tarjanDFS(v Vertex) {
	if color := kg.visited[v]; color == GREY || color == BLACK {
		return
	}

	kg.visited[v] = GREY

	for _, vertex := range kg.graph[v] {
		if vertex == -1 {
			continue
		}
		kg.tarjanDFS(Vertex(vertex))
	}

	kg.visited[v] = BLACK
	kg.vertices.Push(v)
}
