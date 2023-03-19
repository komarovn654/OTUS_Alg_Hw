package graph

type AdjacentVertex struct {
	Num    int
	Weight int
}

type AdjacentVector []AdjacentVertex

type Graph struct {
	vertices []AdjacentVector
}

type Edge struct {
	Weight  int
	Vertex1 int
	Vertex2 int
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
	visited := make(map[int]bool)

	for vertex1, adjVector := range g.vertices {
		for _, adjVertex := range adjVector {
			if _, ok := visited[adjVertex.Num]; !ok {
				edges = append(edges, Edge{Weight: adjVertex.Weight, Vertex1: vertex1, Vertex2: adjVertex.Num})
			}
		}
		visited[vertex1] = true
	}

	return edges
}
