package main

import (
	"fmt"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/hw04arrays"
)

func main() {
	na := hw04arrays.InitDynamicArray(1)

	timer := time.Now()
	for i := 0; i < 100000; i++ {
		na.Add(hw04arrays.Item(i))
	}
	fmt.Println(time.Since(timer))
}
