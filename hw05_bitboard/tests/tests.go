package main

import (
	"errors"
	"log"

	"github.com/komarovn654/OTUS_Alg_Hw/hw05bitboard"
)

var (
	ErrMaskMismatch  = errors.New("mismatch mask")
	ErrCountShift    = errors.New("count error, shift method")
	ErrCountSubtract = errors.New("count error, subtract method")
	ErrCountCache    = errors.New("count error, cache method")
)

type testCase struct {
	in     uint64 // input value
	expCnt uint64 // expected count
	expPos uint64 // expected positions
}

type testResult struct {
	tc  testCase
	res bool
	err error
}

func runCountTest(mask uint64, t testCase) error {
	cnt := hw05bitboard.BitsCountShift(mask)
	log.Printf("count shift: %v", cnt)
	if t.expCnt != cnt {
		return ErrCountShift
	}

	cnt = hw05bitboard.BitsCountSubtract(mask)
	log.Printf("count subtract: %v", cnt)
	if t.expCnt != cnt {
		return ErrCountSubtract
	}

	cnt = hw05bitboard.BitsCountCache(mask, &cache)
	log.Printf("count cache: %v", cnt)
	if t.expCnt != cnt {
		return ErrCountCache
	}

	return nil
}

func runTest(tests []testCase, cp hw05bitboard.ChessPiece) []testResult {
	tr := make([]testResult, 0)
	for i, t := range tests {
		log.Printf("test case: %+v", t)
		tr = append(tr, testResult{})

		mask := cp.GetMoves(t.in)
		log.Printf("positions: %v", mask)
		if t.expPos != mask {
			tr[i] = testResult{tc: t, res: false, err: ErrMaskMismatch}
			return tr
		}

		if err := runCountTest(mask, t); err != nil {
			tr[i] = testResult{tc: t, res: false, err: err}
			return tr
		}

		tr[i] = testResult{tc: t, res: true, err: nil}
		log.Printf("Result: %v\n\n", true)
	}

	return tr
}

func kingTest(dir string) ([]testResult, error) {
	log.Println("king test")
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.King{}), nil
}

func knightTest(dir string) ([]testResult, error) {
	log.Println("knight test")
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.Knight{}), nil
}

func rookTest(dir string) ([]testResult, error) {
	log.Println("rook test")
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.Rook{}), nil
}

func bishopTest(dir string) ([]testResult, error) {
	log.Println("bishop test")
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.Bishop{}), nil
}

func queenTest(dir string) ([]testResult, error) {
	log.Println("qeen test")
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.Queen{}), nil
}
