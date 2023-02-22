package hw14_iddfs

import "fmt"

type Vertex int
type Vertices []Vertex
type Graph struct {
	vertices []Vertices // Vertices nums begins with 0
	visited  []bool
}

// depth-limited search
func (g *Graph) DLS() {
	g.visited = make([]bool, len(g.vertices))

	g.dls(0, 2)
}

func (g *Graph) dls(vertex Vertex, limit int) {
	if g.visited[vertex] {
		return
	}
	fmt.Println(vertex)
	g.visited[vertex] = true

	for _, v := range g.vertices[vertex] {
		g.dls(v)
	}
}
