package bridges

import "fmt"

type Vertex int
type Bridge struct {
	from Vertex
	to   Vertex
}
type Graph struct {
	graph   [][]Vertex
	visited map[Vertex]bool
	tin     map[Vertex]int
	tup     map[Vertex]int
	bridges []Bridge
}

func InitGraph(g *[][]Vertex) Graph {
	return Graph{
		graph:   *g,
		visited: make(map[Vertex]bool),
		tin:     make(map[Vertex]int),
		tup:     make(map[Vertex]int),
		bridges: make([]Bridge, 0),
	}
}

func (g *Graph) FindBridges() []Bridge {
	g.DFS()
	return g.bridges
}

func (g *Graph) DFS() {
	if g.graph == nil || g.visited == nil {
		return
	}

	g.dfs(0, -1, 0)
}

func (g *Graph) dfs(v Vertex, parent Vertex, time int) int {
	fmt.Printf("%v tin = %v\n", v, g.tin[v])
	time++
	g.tin[v] = time
	g.tup[v] = time
	g.visited[v] = true

	for _, vertex := range g.graph[v] {
		if vertex == parent {
			continue
		}
		if g.visited[vertex] {
			// vertex is reverse edge from v
			g.tup[v] = MinInt(g.tin[vertex], g.tup[v])
			continue
		}
		g.dfs(vertex, v, time)
		g.tup[v] = MinInt(g.tup[v], g.tup[vertex])
		if g.tin[v] < g.tup[vertex] {
			g.bridges = append(g.bridges, Bridge{v, vertex})
		}
	}

	return g.tup[v]
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
