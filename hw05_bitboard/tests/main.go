package main

import (
	"fmt"
	"log"
	"os"

	"github.com/komarovn654/OTUS_Alg_Hw/hw05bitboard"
)

var (
	cache = hw05bitboard.Cache{}
)

type testCase struct {
	name   string // ChessPiece name
	in     uint64 // input value
	expCnt uint64 // expected count
	expPos uint64 // expected positions
}

type calcResult struct {
	mask uint64
	bits uint64
	res  bool
}

type testResult struct {
	tc  testCase
	res calcResult
	err error
}

func printResult(filePath string, results []testResult) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, tr := range results {
		str := fmt.Sprintf(
			"%v, Position: %v  \n"+
				"Expected Mask - Calculated Mask: %v - %v\n"+
				"Expected Bits - Calculated Bits: %v - %v\n"+
				"Result: %v  \n",
			tr.tc.name, tr.tc.in,
			tr.tc.expPos, tr.res.mask,
			tr.tc.expCnt, tr.res.bits,
			tr.res.res,
		)
		f.WriteString(str)
	}
}

func init() {
	cache.Init()
}

func main() {
	result, err := runTests()
	if err != nil {
		log.Fatal(err)
	}

	printResult("README.md", result)
}
