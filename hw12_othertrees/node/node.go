package node

import (
	"math/rand"
	"time"
)

type Node interface {
	GetKey() int
	SetKey(int)
	GetWeight() int
	SetWeight(int)
}

type node struct {
	key    int
	weight int
}

// Generate random nodes with unique keys
func GenerateNodes(length int, keyRange int, weightRange int) []Node {
	Nodes := make([]Node, length)
	nodes := make([]node, length)
	isExist := make(map[int]bool)

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; len(isExist) < length; {
		key := rnd.Intn(keyRange)
		if !isExist[key] {
			isExist[key] = true

			nodes[i].key = key
			nodes[i].weight = rnd.Intn(keyRange)
			Nodes[i] = &nodes[i]
			i++
		}
	}

	return Nodes
}

func (v *node) GetKey() int {
	return v.key
}

func (v *node) GetWeight() int {
	return v.weight
}

func (v *node) SetKey(key int) {
	v.key = key
}

func (v *node) SetWeight(weight int) {
	v.weight = weight
}
