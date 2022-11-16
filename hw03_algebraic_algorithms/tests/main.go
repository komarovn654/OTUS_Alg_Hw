package main

import (
	"errors"
	"flag"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/komarovn654/OTUS_Alg_Hw/hw03alg"
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

func readLine(filePath string) (string, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return strings.Split(string(f), "\n")[0], nil
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

func fibonacci(alg string, n int) (*big.Int, error) {
	switch alg {
	case "recursive":
		return hw03alg.FibRecursive(n), nil
	case "iterative":
		return hw03alg.FibIterative(n), nil
	case "golden ratio":
		return hw03alg.FibGolden(n), nil
	case "matrix":
		return hw03alg.FibMatrix(n), nil
	}
	return nil, ErrUnsupAlg
}

func init() {
	flag.StringVar(&in, "in", "", "file with input parameters")
	flag.StringVar(&out, "out", "", "file with parameters for comparison")
	flag.StringVar(&task, "task", "", "task: fibonacci, prime or power")
	flag.StringVar(&alg, "alg", "", "algorithm name")
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
