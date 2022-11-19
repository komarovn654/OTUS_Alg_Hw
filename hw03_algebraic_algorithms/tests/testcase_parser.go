package main

import (
	"bufio"
	"math/big"
	"os"
	"strconv"
)

// Parse fibonacci number
func parseFibN(filePath string) (intN int, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	fscan := bufio.NewScanner(f)

	if !fscan.Scan() {
		return
	}
	intN, err = strconv.Atoi(fscan.Text())
	if err != nil {
		return
	}

	return
}

// Parse base and degree
func parsePwr(filePath string) (base float64, exp int64, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	fscan := bufio.NewScanner(f)

	if !fscan.Scan() {
		return
	}
	base, err = strconv.ParseFloat(fscan.Text(), 64)
	if err != nil {
		return
	}

	if !fscan.Scan() {
		return
	}
	exp, err = strconv.ParseInt(fscan.Text(), 10, 64)
	if err != nil {
		return
	}

	return
}

// Parse expected result big Int
func parseResBInt(filePath string) (bInt *big.Int, err error) {
	bInt = big.NewInt(0)
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	fscan := bufio.NewScanner(f)

	if !fscan.Scan() {
		return
	}
	if bInt, ok := bInt.SetString(fscan.Text(), 10); !ok {
		return bInt, err
	}

	return
}

// Parse expected result float
func parseResFloat(filePath string) (res float32, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	fscan := bufio.NewScanner(f)
	if !fscan.Scan() {
		return
	}
	res64, err := strconv.ParseFloat(fscan.Text(), 64)
	if err != nil {
		return
	}
	return float32(res64), err
}

// Parse int64 from first string
func parseInt(filePath string) (res int64, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}

	fscan := bufio.NewScanner(f)
	if !fscan.Scan() {
		return
	}

	res, err = strconv.ParseInt(fscan.Text(), 10, 64)
	if err != nil {
		return
	}
	return
}
