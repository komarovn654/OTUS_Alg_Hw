package main

import (
	"fmt"
	"time"

	alg1 "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/alg1"
	alg2 "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/alg2"
	node "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

func main() {
	vertices := node.GenerateNodes(1_000_000, 1_000_000, 1_000_000)

	t := time.Now()
	bst1 := alg1.BuildBST(vertices)
	fmt.Println("Alg1 1_000_000 nodes: ", time.Since(t))
	fmt.Println("BST1 is valid: ", bst1.IsValid())

	t = time.Now()
	bst2 := alg2.BuildBST(vertices)
	fmt.Println("Alg2 1_000_000 nodes: ", time.Since(t))
	fmt.Println("BST2 is valid: ", bst2.IsValid())
}
