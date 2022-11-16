package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/hw03alg"
)

var (
	ErrUnsupAlg  = errors.New("unsupported algorithm")
	ErrUnsupTask = errors.New("unsupported task")
	ErrBigintStr = errors.New("string to bigint cast error")

	in   string
	out  string
	task string
	alg  string
)

func readLine(filePath string) (string, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return strings.Split(string(f), "\n")[0], nil
}

func parseN(filePath string) (int, error) {
	strN, err := readLine(filePath)
	if err != nil {
		return 0, err
	}
	intN, err := strconv.Atoi(strN)
	if err != nil {
		return 0, err
	}
	return intN, nil
}

func parseResult(filePath string) (*big.Int, error) {
	strRes, err := readLine(filePath)
	if err != nil {
		return big.NewInt(0), err
	}
	bintRes := big.NewInt(0)
	if _, ok := bintRes.SetString(strRes, 10); !ok {
		return big.NewInt(0), ErrBigintStr
	}

	return bintRes, nil
}

func fibonacci(alg string, n int) (*big.Int, error) {
	switch alg {
	case "recursive":
		return hw03alg.FibRecursive(n), nil
	case "iterative":
		return hw03alg.FibIterative(n), nil
	case "goldenratio":
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
	resCmp, err := parseResult(out)
	if err != nil {
		log.Fatal(err)
	}

	rdy := make(chan struct{})

	res := big.NewInt(0)
	timer := time.Now()
	go func() {
		switch task {
		case "fibonacci":
			res, err = fibonacci(alg, n)
		case "prime":
		case "power":
		default:
			log.Fatal(ErrUnsupTask)
		}
		if err != nil {
			log.Fatal(err)
		}
		rdy <- struct{}{}
	}()

	select {
	case <-rdy:
		if resCmp.Cmp(res) == 0 {
			fmt.Printf(" %v, EXECUTION TIME: %v\n", "PASS", time.Since(timer))
			return
		}
		fmt.Println("FAIL")
	case <-time.After(60 * time.Second):
		fmt.Println("TIMEOUT")
	}
}
