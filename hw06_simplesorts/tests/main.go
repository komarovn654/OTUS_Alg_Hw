package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	hw06simplesorts "github.com/komarovn654/OTUS_Alg_Hw/hw06_simplesorts"
)

var sortTimeout = time.Second * 120

func randArray() []hw06simplesorts.Item {
	array := make([]hw06simplesorts.Item, 10)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range array {
		array[i] = hw06simplesorts.Item(r1.Int63n(100))
	}
	return array
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), sortTimeout)
	defer cancel()

	ar := randArray()
	res := hw06simplesorts.SortArray(ctx, &ar, hw06simplesorts.BubbleSort)
	fmt.Printf("%v\n", res.Time.Time)

	ar = randArray()
	res = hw06simplesorts.SortArray(ctx, &ar, hw06simplesorts.InsertionSort)
	fmt.Printf("%v\n", res.Time.Time)

	ar = randArray()
	fmt.Println(ar)
	res = hw06simplesorts.SortArray(ctx, &ar, hw06simplesorts.ShellSort)
	fmt.Println(ar)
	fmt.Printf("%v\n", res.Time.Time)
}
