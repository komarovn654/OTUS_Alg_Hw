package iddfs

const (
	STATE_VISITED_ITEM = 1
	STATE_EDGE_ITEM    = 2

	ITERATION_STEP = 1

	START_VERTEX = 0
)

type Vertex int
type Vertices []Vertex
type Graph struct {
	vertices      []Vertices // Vertices nums begins with 0
	visited       map[Vertex]int
	startVertices []Vertex
}

// iterative deepening depth-first search
func (g *Graph) IDDFS(target Vertex) bool {
	g.visited = make(map[Vertex]int, len(g.vertices))
	g.startVertices = []Vertex{START_VERTEX}

	for limit := 0; limit < len(g.vertices); limit += ITERATION_STEP {
		for _, start := range g.startVertices {
			if g.dls(start, target, ITERATION_STEP) {
				return true
			}
		}

		// get edge items and make them the starting items in next iteration
		g.startVertices = make([]Vertex, 0)
		for vertex, state := range g.visited {
			if state == STATE_EDGE_ITEM {
				g.startVertices = append(g.startVertices, vertex)
				g.visited[vertex] = 0
			}
		}
	}

	return false
}

// depth-limited search
func (g *Graph) dls(start Vertex, target Vertex, depth int) (result bool) {
	if g.visited[start] == STATE_VISITED_ITEM || start == target {
		return start == target
	}

	if depth == 0 {
		g.visited[start] = STATE_EDGE_ITEM
		return start == target
	}

	g.visited[start] = STATE_VISITED_ITEM

	for _, v := range g.vertices[start] {
		if result = g.dls(v, target, depth-1); result {
			break
		}
	}

	return result
}
