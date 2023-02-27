package demukron

import (
	hw15 "github.com/komarovn654/OTUS_Alg_Hw/hw15_demukron"
)

type DemukronGraph struct {
	graph     hw15.Graph
	degrees   []int
	processed int // vertices processed
}

func InitDemukron(g *hw15.Graph) DemukronGraph {
	return DemukronGraph{
		graph:   *g,
		degrees: g.GetDegrees(),
	}
}

// Return nil if the graph is not initialized or a loop is detected
func (g *DemukronGraph) Sort() (levels []hw15.Vertices) {
	if g.degrees == nil || g.graph == nil {
		return nil
	}

	levels = make([]hw15.Vertices, 0)

	for g.processed != len(g.graph) {
		levels = append(levels, g.nextLevel())
	}

	return
}

func (g *DemukronGraph) nextLevel() (level hw15.Vertices) {
	level = make(hw15.Vertices, 0)

	// build a level of vertices from zero inputs
	for vertex, degree := range g.degrees {
		if degree == 0 {
			level = append(level, hw15.Vertex(vertex))
		}
	}

	// modify slice of degrees
	for _, levelVertex := range level {
		g.processed++
		g.degrees[levelVertex]-- // set negatiaon value in zero inputs vertex
		for _, input := range g.graph[levelVertex] {
			g.degrees[input]--
		}
	}

	return level
}
