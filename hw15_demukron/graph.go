package graph

type Vertex int
type Vertices []Vertex
type Graph []Vertices

// Return the degrees of vertices.
// Return nil if the graph is not initialized or a loop is detected
func (g *Graph) GetDegrees() []int {
	if g == nil {
		return nil
	}

	degrees := make([]int, len(*g))

	for i := range *g {
		for _, vertex := range (*g)[i] {
			if vertex == Vertex(i) {
				return nil // loop
			}
			degrees[vertex]++
		}
	}

	return degrees
}
