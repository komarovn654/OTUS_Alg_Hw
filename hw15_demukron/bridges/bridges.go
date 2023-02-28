package bridges

import (
	"fmt"
)

type Vertex struct {
	num  int
	time int
	min  int // min time
}

type Vertices []Vertex
type Graph struct {
	graph   []Vertices
	visited map[int]bool
}

func (g *Graph) Bridges() {

}

func (g *Graph) DFS() {
	if g.graph == nil || g.visited == nil {
		return
	}

	g.dfs(&g.graph[0][0], 0)
}

func (g *Graph) dfs(v *Vertex, time int) {
	if g.visited[v.num] {
		return
	}

	time++
	v.time += time
	g.visited[v.num] = true
	fmt.Println(v)
	for _, vertex := range g.graph[v.num] {
		g.dfs(&vertex, vertex.time)
	}
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
