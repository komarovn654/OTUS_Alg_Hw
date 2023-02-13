package hw14_kosaraju

type Vertice []int
type Graph []Vertice // Vertices nums begins with 0

// edge

// invert graph
func (g *Graph) reverse() *Graph {
	if len(*g) == 0 {
		return g
	}

	// init empty graph
	invert := make(Graph, len(*g))
	for i := range invert {
		invert[i] = make(Vertice, 0)
	}

	// invert graph
	for i, vertice := range *g {
		for _, value := range vertice {
			if value == -1 {
				continue
			}
			invert[value] = append(invert[value], i)
		}
	}

	return &invert
}
