package graph

import "math"

type AdjacentVertex struct {
	Dst    int
	Weight int
}

type AdjacentVector []AdjacentVertex

type Graph struct {
	vertices []AdjacentVector
}

type Edge struct {
	Weight int
	Src    int
	Dst    int
}

func Init() Graph {
	return Graph{
		vertices: make([]AdjacentVector, 0),
	}
}

func (g *Graph) AddAdjacentVector(vector AdjacentVector) {
	g.vertices = append(g.vertices, vector)
}

func (g *Graph) GetVertices() []int {
	vertices := make([]int, len(g.vertices))

	for i := range g.vertices {
		vertices[i] = i
	}

	return vertices
}

func (g *Graph) GetAdjacentVectors() []AdjacentVector {
	return g.vertices
}

func (g *Graph) GetEdges() []Edge {
	edges := make([]Edge, 0)

	for vertex1, adjVector := range g.vertices {
		for _, adjVertex := range adjVector {
			edges = append(edges, Edge{Weight: adjVertex.Weight, Src: vertex1, Dst: adjVertex.Dst})
		}
	}

	return edges
}

func (g *Graph) CheapestEdge(sourceVertex int, visited map[Edge]bool) Edge {
	min := Edge{Weight: math.MaxInt, Src: sourceVertex, Dst: -1}
	for _, vertex := range g.vertices[sourceVertex] {
		directEdge := Edge{Weight: vertex.Weight, Src: sourceVertex, Dst: vertex.Dst}

		if vertex.Weight < min.Weight && !directEdge.isVisited(visited) {
			min = directEdge
		}
	}

	return min
}

func (this *Edge) isVisited(visited map[Edge]bool) bool {
	reverseEdge := Edge{Weight: this.Weight, Src: this.Dst, Dst: this.Src}

	if visited[*this] || visited[reverseEdge] {
		return true
	}

	return false
}
