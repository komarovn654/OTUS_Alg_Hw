package hw12othertrees

type nodeItem struct {
	key   int
	value []int
}

type splayNode struct {
	item  nodeItem
	left  *splayNode
	right *splayNode
}

func (sn *splayNode) search(key int) *splayNode {

	return nil
}
