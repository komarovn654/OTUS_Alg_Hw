package hw14_kosaraju

const (
	WHITE = iota
	GREY
	BLACK
)

type vertices map[Vertex]int // { vertex: tarjan's color}

type tarjanGraph struct {
	g       Graph
	visited vertices
}

func (tg *tarjanGraph) DFS(start Vertex) vertices {
	if tg == nil {
		return vertices{}
	}

	if len(tg.g) == 0 {
		return vertices{}
	}

	tg.visited = make(vertices)

	tg.tarjanDFS(start)
	return tg.visited
}

func (tg *tarjanGraph) tarjanDFS(v Vertex) {
	if color := tg.visited[v]; color == GREY || color == BLACK {
		return
	}

	tg.visited[v] = GREY

	for _, vertex := range tg.g[v] {
		if vertex == -1 {
			continue
		}
		tg.tarjanDFS(Vertex(vertex))
	}

	tg.visited[v] = BLACK
}
