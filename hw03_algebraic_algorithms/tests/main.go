package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/komarovn654/OTUS_Alg_Hw/hw03alg"
)

var (
	ErrUnsupAlg      = errors.New("unsupported algorithm")
	ErrUnsupTask     = errors.New("unsupported task")
	ErrBigintStr     = errors.New("string to bigint cast error")
	ErrStackOverflow = errors.New("stack overflow")
)

type testResult struct {
	res bool
	dur time.Duration
	err error
}

func fibonacci(alg string, in string, out string) testResult {
	var res *big.Int

	n, err := parseFibN(in)
	if err != nil {
		return testResult{false, 0, err}
	}

	timer := time.Now()
	switch alg {
	case "recursive":
		if n == 10_000_000 {
			return testResult{false, 0, err}
		}
		res = hw03alg.FibRecursive(n)
	case "iterative":
		res = hw03alg.FibIterative(n)
	case "goldenratio":
		res = hw03alg.FibGolden(n)
	case "matrix":
		res = hw03alg.FibMatrix(n)
	default:
		return testResult{false, 0, ErrUnsupAlg}
	}
	runtime := time.Since(timer)

	resCmp, err := parseResBInt(out)
	if err != nil {
		return testResult{false, 0, err}
	}

	if resCmp.Cmp(res) != 0 {
		return testResult{false, runtime, nil}
	}
	return testResult{true, runtime, nil}
}

func power(alg string, in string, out string) testResult {
	var res float32

	base, degree, err := parsePwr(in)
	if err != nil {
		return testResult{false, 0, err}
	}

	timer := time.Now()
	switch alg {
	case "iterative":
		res = hw03alg.PwrIterative(base, degree)
	case "sqrmultiply":
		res = hw03alg.PwrSqrMultiply(base, degree)
	case "binary":
		res = hw03alg.PwrBinary(base, degree)
	default:
		return testResult{false, 0, ErrUnsupAlg}
	}
	runtime := time.Since(timer)

	resCmp, err := parseResFloat(out)
	if err != nil {
		return testResult{false, 0, err}
	}

	res = (res * 100_000_000_000) / 100_000_000_000
	if resCmp != res {
		return testResult{false, runtime, nil}
	}
	return testResult{true, runtime, nil}
}

func prime(alg string, in string, out string) testResult {
	var res int64

	n, err := parseInt(in)
	if err != nil {
		return testResult{false, 0, err}
	}

	timer := time.Now()
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
		return testResult{false, 0, ErrUnsupAlg}
	}
	runtime := time.Since(timer)

	resCmp, err := parseInt(out)
	if err != nil {
		return testResult{false, 0, err}
	}

	if resCmp != res {
		return testResult{false, runtime, nil}
	}
	return testResult{true, runtime, nil}
}

func worker(task string, alg string, in string, out string) <-chan testResult {
	tr := testResult{}
	done := make(chan testResult)

	go func() {
		switch task {
		case "fibonacci":
			tr = fibonacci(alg, in, out)
		case "prime":
			tr = prime(alg, in, out)
		case "power":
			tr = power(alg, in, out)
		default:
			tr = testResult{false, 0, ErrUnsupTask}
		}
		done <- tr
	}()

	return done
}

func saveResult(filePath string, result testResult, timeout bool) error {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	if result.res {
		_, err = f.WriteString(fmt.Sprintf("| PASS_%v*", result.dur))
		return err
	}

	if timeout {
		_, err = f.WriteString(fmt.Sprintf("| %v*", "TIMEOUT"))
		return err
	}

	_, err = f.WriteString(fmt.Sprintf("| FAIL %v*", result.dur))
	return err
}

func main() {
	if len(os.Args) < 5 {
		log.Fatalf("Five arguments are expected: \n%v\n%v\n%v\n%v\n%v\n",
			"1. file with input parameters",
			"2. file with parameters for comparison",
			"3. task: fibonacci, prime or power",
			"4. algorithm name",
			"5. file for saving the result",
		)
	}

	result := worker(os.Args[3], os.Args[4], os.Args[1], os.Args[2])

	select {
	case res := <-result:
		if res.err != nil && res.err != ErrStackOverflow {
			log.Fatal(res.err)
		}
		if err := saveResult(os.Args[5], res, false); err != nil {
			log.Fatal(err)
		}
	case <-time.After(60 * time.Second):
		if err := saveResult(os.Args[5], testResult{}, true); err != nil {
			log.Fatal(err)
		}
	}
}
