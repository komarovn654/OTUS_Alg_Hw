package unionfind

// pair vertex->root
type pair map[int]int

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
func (u *UnionFind) AddPair(vertex int, root int) {
	u.vertices[vertex] = root
}

func (u *UnionFind) GetRoots() []int {
	roots := make([]int, 0)
	added := make(map[int]bool)
	for _, root := range u.vertices {
		if exist, ok := added[root]; ok && exist {
			continue
		}
		added[root] = true
		roots = append(roots, root)
	}

	return roots
}

func (u *UnionFind) GetPairs() pair {
	return u.vertices
}

func (u *UnionFind) Find(vertex int) (root int) {
	root = u.find(vertex)
	if root != -1 {
		u.vertices[vertex] = root
	}

	return root
}

func (u *UnionFind) find(vertex int) (root int) {
	root, ok := u.vertices[vertex]
	if !ok {
		return -1
	}
	if root == vertex {
		return root
	}

	return u.find(root)
}

func (u *UnionFind) Union(v1 int, v2 int) (root int) {
	root = u.Find(v1)

	u.vertices[v2] = root
	return root
}
