package main

import (
	"errors"

	"github.com/komarovn654/OTUS_Alg_Hw/hw05bitboard"
)

var (
	ErrMaskMismatch  = errors.New("mismatch mask")
	ErrCountShift    = errors.New("count error, shift method")
	ErrCountSubtract = errors.New("count error, subtract method")
	ErrCountCache    = errors.New("count error, cache method")
)

func runCountTest(mask uint64, t testCase) (uint64, error) {
	cnt := hw05bitboard.BitsCountShift(mask)
	if t.expCnt != cnt {
		return 0, ErrCountShift
	}

	cnt = hw05bitboard.BitsCountSubtract(mask)
	if t.expCnt != cnt {
		return 0, ErrCountSubtract
	}

	cnt = hw05bitboard.BitsCountCache(mask, &cache)
	if t.expCnt != cnt {
		return 0, ErrCountCache
	}

	return cnt, nil
}

func runTest(tests []testCase, cp hw05bitboard.ChessPiece) []testResult {
	tr := make([]testResult, 0)
	for i, t := range tests {
		tr = append(tr, testResult{})

		mask := cp.GetMoves(t.in)
		if t.expPos != mask {
			tr[i] = testResult{tc: t, res: calcResult{}, err: ErrMaskMismatch}
			return tr
		}

		cnt, err := runCountTest(mask, t)
		if err != nil {
			tr[i] = testResult{tc: t, res: calcResult{}, err: err}
			return tr
		}

		tr[i] = testResult{tc: t, res: calcResult{mask, cnt, true}, err: nil}
	}

	return tr
}

func kingTest(dir string) ([]testResult, error) {
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.King{}), nil
}

func knightTest(dir string) ([]testResult, error) {
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.Knight{}), nil
}

func rookTest(dir string) ([]testResult, error) {
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.Rook{}), nil
}

func bishopTest(dir string) ([]testResult, error) {
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.Bishop{}), nil
}

func queenTest(dir string) ([]testResult, error) {
	tc, err := parseTestCases(dir)
	if err != nil {
		return nil, err // TODO: separate errors
	}

	return runTest(tc, &hw05bitboard.Queen{}), nil
}

func runTests() ([]testResult, error) {
	results := make([]testResult, 0)

	tr, err := kingTest("1.Bitboard - Король")
	if err != nil {
		return nil, err
	}
	results = append(results, tr...)

	tr, err = knightTest("2.Bitboard - Конь")
	if err != nil {
		return nil, err
	}
	results = append(results, tr...)

	tr, err = rookTest("3.Bitboard - Ладья")
	if err != nil {
		return nil, err
	}
	results = append(results, tr...)

	tr, err = bishopTest("4.Bitboard - Слон")
	if err != nil {
		return nil, err
	}
	results = append(results, tr...)

	tr, err = queenTest("5.Bitboard - Ферзь")
	if err != nil {
		return nil, err
	}
	results = append(results, tr...)

	return results, err
}
