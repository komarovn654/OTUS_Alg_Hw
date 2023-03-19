package unionfind

// pair vertex->root
type pair map[any]any

type UnionFind struct {
	vertices pair
}

func Init(vertices []int) UnionFind {
	pairs := make(pair)
	for _, vertex := range vertices {
		pairs[vertex] = vertex
	}

	return UnionFind{
		vertices: pairs,
	}
}

// Add pair to UnionFind. If a vertex already exists, replaces it.
func (u *UnionFind) AddPair(vertex any, root any) {
	u.vertices[vertex] = root
}

func (u *UnionFind) GetPairs() pair {
	return u.vertices
}

func (u *UnionFind) Find(vertex any) (root any) {
	root = u.find(vertex)
	if root != nil {
		u.vertices[vertex] = root
	}

	return root
}

func (u *UnionFind) find(vertex any) (root any) {
	root, ok := u.vertices[vertex]
	if !ok {
		return nil
	}
	if root == vertex {
		return root
	}

	return u.find(root)
}

func (u *UnionFind) Union(v1 any, v2 any) (root any) {
	root = u.Find(v1)

	u.vertices[v2] = root
	return root
}
