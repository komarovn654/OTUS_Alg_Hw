package main

import (
	"errors"
	"flag"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	ErrUnsupAlg  = errors.New("unsupported algorithm")
	ErrUnsupTask = errors.New("unsupported task")

	in   string
	out  string
	task string
	alg  string
)

type testStruct struct {
	n      int
	expect *big.Int
}

func parseN(filePath string) (int, error) {
	strN, err := readLine(in)
	if err != nil {
		return 0, err
	}
	intN, err := strconv.Atoi(strN)
	if err != nil {
		return 0, err
	}
	return intN, nil
}

func readLine(filePath string) (string, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return strings.Split(string(f), "\n")[0], nil
}

func prepareTest(testNum int) (testStruct, error) {
	i, _ := readLine(in)
	o, _ := readLine(out)

	bio := big.NewInt(0)
	expect, _ := bio.SetString(o, 10)

	return testStruct{n, expect}, nil
}

func init() {
	flag.StringVar(&in, "in", "", "file with input parameters")
	flag.StringVar(&out, "out", "", "file with parameters for comparison")
	flag.StringVar(&task, "task", "", "task: fibonacci, prime or power")
	flag.StringVar(&alg, "alg", "", "algorithm name")
}

func fibonacci(alg string, n int) (*big.Int, error) {
	switch alg {
	case "recursive":
		return hw03alg.FibRecursive(n)
	case "iterative":
	case "golden ratio":
	case "matrix":
	}
	return nil, ErrUnsupAlg
}

func main() {
	flag.Parse()
	n, err := parseN(in)
	if err != nil {
		log.Fatal(err)
	}

	switch task {
	case "fibonacci":
		fibonacci(alg, n)
	case "prime":
	case "power":
	default:
		log.Fatal(ErrUnsupTask)
	}
}
