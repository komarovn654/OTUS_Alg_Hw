package ucs

import (
	"sort"
)

type Vertex struct {
	num  int
	cost int
}
type Vertices []Vertex
type Graph struct {
	vertices []Vertices // Vertices nums begins with 0
	queue    List
	visited  map[int]bool
}

func (g *Graph) UCS(target int) (sumCost int, exist bool) {
	g.queue = NewQueue()
	g.visited = make(map[int]bool)

	for _, vertex := range g.vertices {
		vertex.sortByCost()
	}

	g.queue.Queue(Vertex{num: 0})

	for {
		vertex, ok := g.queue.DequeueWithCast()
		if !ok {
			break // queue is empty
		}
		sumCost = vertex.cost

		g.visited[vertex.num] = true
		if vertex.num == target {
			return sumCost, true
		}

		for _, v := range g.vertices[vertex.num] {
			if !g.visited[v.num] {
				v.cost += sumCost
				g.queue.Queue(v)
			}
		}
	}

	return sumCost, false
}

func (l *List) DequeueWithCast() (Vertex, bool) {
	queueItem, ok := l.Dequeue()
	if !ok {
		return Vertex{}, false // queue is empty
	}
	vertex, ok := queueItem.(Vertex)
	if !ok {
		panic("type cast error")
	}

	return vertex, true
}

func (v *Vertices) sortByCost() {
	sort.Slice(*v, func(i, j int) bool {
		return (*v)[i].cost < (*v)[j].cost
	})
}
