package main

import (
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Parse fibonacci number
func parseFibN(filePath string) (int, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	intN, err := strconv.Atoi(strings.Split(string(f), "\n")[0])
	if err != nil {
		return 0, err
	}

	return intN, nil
}

// Parse base and degree
func parsePwr(filePath string) (float64, int64, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return 0, 0, err
	}

	fileStrings := strings.Split(string(f), "\r\n") // ???

	base, err := strconv.ParseFloat(fileStrings[0], 64)
	if err != nil {
		return 0, 0, err
	}

	degree, err := strconv.ParseInt(fileStrings[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return base, degree, nil
}

// Parse expected result big Int
func parseResBInt(filePath string) (*big.Int, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return big.NewInt(0), err
	}

	bintRes := big.NewInt(0)
	if _, ok := bintRes.SetString(strings.Split(string(f), "\n")[0], 10); !ok {
		return big.NewInt(0), ErrBigintStr
	}

	return bintRes, nil
}

// Parse expected result float
func parseResFloat(filePath string) (float32, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	res, err := strconv.ParseFloat(strings.Split(string(f), "\n")[0], 64)
	if err != nil {
		return 0, err
	}

	return float32(res), nil
}

// Parse int64 from first string
func parseInt(filePath string) (int64, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	res, err := strconv.ParseInt(strings.Split(string(f), "\r\n")[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return res, nil
}
