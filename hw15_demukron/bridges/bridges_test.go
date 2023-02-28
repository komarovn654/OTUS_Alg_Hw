package bridges

import (
	"testing"
)

func TestDFS(t *testing.T) {
	test := []Vertices{
		{Vertex{num: 1}, Vertex{num: 2}},                 // 0
		{Vertex{num: 0}, Vertex{num: 2}},                 // 1
		{Vertex{num: 0}, Vertex{num: 1}, Vertex{num: 3}}, // 2
		{Vertex{num: 2}, Vertex{num: 4}, Vertex{num: 5}}, // 3
		{Vertex{num: 3}, Vertex{num: 5}},                 // 4
		{Vertex{num: 3}, Vertex{num: 4}},                 // 5
	}
	tg := Graph{graph: test, visited: make(map[int]bool)}

	tg.DFS()
}
