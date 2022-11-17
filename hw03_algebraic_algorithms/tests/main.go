package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/big"
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

func fibonacci(alg string, in string, out string) (bool, error) {
	var res *big.Int

	n, err := parseFibN(in)
	if err != nil {
		return false, err
	}

	switch alg {
	case "recursive":
		res = hw03alg.FibRecursive(n)
	case "iterative":
		res = hw03alg.FibIterative(n)
	case "goldenratio":
		res = hw03alg.FibGolden(n)
	case "matrix":
		res = hw03alg.FibMatrix(n)
	default:
		return false, ErrUnsupAlg
	}

	resCmp, err := parseResBInt(out)
	if err != nil {
		return false, err
	}

	if resCmp.Cmp(res) != 0 {
		return false, nil
	}
	return true, nil
}

func power(alg string, in string, out string) (bool, error) {
	var res float32

	base, degree, err := parsePwr(in)
	if err != nil {
		return false, err
	}

	switch alg {
	case "iterative":
		res = hw03alg.PwrIterative(base, degree)
	case "sqrmultiply":
		res = hw03alg.PwrSqrMultiply(base, degree)
	case "binary":
		res = hw03alg.PwrBinary(base, degree)
	default:
		return false, ErrUnsupAlg
	}

	resCmp, err := parseResFloat(out)
	if err != nil {
		return false, err
	}

	res = (res * 100_000_000_000) / 100_000_000_000
	if resCmp != res {
		return false, nil
	}
	return true, nil
}

func prime(alg string, in string, out string) (bool, error) {
	var res int64

	n, err := parseInt(in)
	if err != nil {
		return false, err
	}

	switch alg {
	case "brutforce":
		res = hw03alg.PrimeBruteforce(n)
	case "brutforceopt":
		res = hw03alg.PrimeBFOpt(n)
	case "erat":
		res = hw03alg.PrimeErat(n)
	case "eratmem":
		res = hw03alg.PrimeEratMemOpt(n)
	case "eratopt":
		res = hw03alg.PrimeEratOpt(n)
	default:
		return false, ErrUnsupAlg
	}

	resCmp, err := parseInt(out)
	if err != nil {
		return false, err
	}

	if resCmp != res {
		return false, nil
	}

	return true, nil
}

func init() {
	flag.StringVar(&in, "in", "", "file with input parameters")
	flag.StringVar(&out, "out", "", "file with parameters for comparison")
	flag.StringVar(&task, "task", "", "task: fibonacci, prime or power")
	flag.StringVar(&alg, "alg", "", "algorithm name")
}

func main() {
	flag.Parse()
	rdy := make(chan struct{})
	timer := time.Now()
	var res bool
	var err error

	go func() {
		switch task {
		case "fibonacci":
			res, err = fibonacci(alg, in, out)
		case "prime":
			res, err = prime(alg, in, out)
		case "power":
			res, err = power(alg, in, out)
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
		if res {
			fmt.Printf("PASS, EXECUTION TIME: %v\n", time.Since(timer))
			return
		}
		fmt.Println("FAIL")
	case <-time.After(60 * time.Second):
		fmt.Println("TIMEOUT")
	}
}
