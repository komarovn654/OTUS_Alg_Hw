package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrArrayIsUnsorted = errors.New("array is not sorted")
)

const (
	testsDir  = "sorting-tests/0.random,sorting-tests/1.digits,sorting-tests/2.sorted,sorting-tests/3.revers"
	sizeCount = 8
)

func main() {

	tr, err := RunTest()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tr)

}
