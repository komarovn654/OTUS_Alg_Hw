package main

import (
	"fmt"

	node "github.com/komarovn654/OTUS_Alg_Hw/hw12_othertrees/node"
)

func main() {
	vertices := node.GenerateNodes(10, 100, 100)
	fmt.Println(vertices)
}
