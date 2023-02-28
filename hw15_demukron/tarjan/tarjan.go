package tarjan

import (
	hw15 "github.com/komarovn654/OTUS_Alg_Hw/hw15_demukron"
)

const (
	WHITE = iota
	GREY
	BLACK
)

type tarjanVerts map[hw15.Vertex]int // { vertex: tarjan's color}

type TarjanGraph struct {
	graph    hw15.Graph
	visited  tarjanVerts
	vertices hw15.Vertices
}

func InitTarjan(g *hw15.Graph) TarjanGraph {
	return TarjanGraph{
		graph:    *g,
		visited:  make(tarjanVerts),
		vertices: make(hw15.Vertices, 0),
	}
}

func (kg *TarjanGraph) DFS(start hw15.Vertex) hw15.Vertices {
	if kg == nil {
		return kg.vertices
	}

	if len(kg.graph) == 0 {
		return kg.vertices
	}

	kg.dfs(start)
	return kg.vertices
}

func (kg *TarjanGraph) dfs(v hw15.Vertex) {
	if color := kg.visited[v]; color == GREY || color == BLACK {
		return
	}

	kg.visited[v] = GREY

	for _, vertex := range kg.graph[v] {
		if vertex == -1 {
			continue
		}
		kg.dfs(hw15.Vertex(vertex))
	}

	kg.visited[v] = BLACK
	kg.vertices = append(kg.vertices, v)
}
