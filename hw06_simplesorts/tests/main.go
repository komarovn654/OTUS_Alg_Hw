package main

import (
	"context"
	"fmt"
	"time"

	hw06simplesorts "github.com/komarovn654/OTUS_Alg_Hw/hw06_simplesorts"
)

var sortTimeout = time.Second * 120

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), sortTimeout)
	defer cancel()

	ar := hw06simplesorts.RandArray(100, 100)
	res := hw06simplesorts.SortArray(ctx, &ar, hw06simplesorts.InsertionSortShift)
	fmt.Printf("%v\n", res.Time.Time)
	fmt.Println(res.Array)
}
