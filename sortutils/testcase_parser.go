package sortutils

import (
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	Name     string
	N        int64
	Array    []Item
	Expected []Item
}

func (tc *testCase) parseN(str *string) (err error) {
	tc.N, err = strconv.ParseInt(*str, 10, 64)
	return err
}

func (tc *testCase) parseArray(str *string) (err error) {
	tc.Array = make([]Item, tc.N)
	for i, item := range strings.Split(*str, " ") {
		value, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return err
		}
		tc.Array[i] = Item(value)
	}
	return nil
}

func (tc *testCase) parseExpected(str *string) (err error) {
	tc.Expected = make([]Item, tc.N)
	for i, item := range strings.Split(*str, " ") {
		value, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return err
		}
		tc.Expected[i] = Item(value)
	}
	return nil
}

func (tc *testCase) ParseTestCase(dir string, inFile string, outFile string) (err error) {
	tc.Name = dir
	inLines, err := readFile(inFile)
	if err != nil {
		return err
	}
	outLines, err := readFile(outFile)
	if err != nil {
		return err
	}

	if err := tc.parseN(&inLines[0]); err != nil {
		return err
	}
	if err := tc.parseArray(&inLines[1]); err != nil {
		return err
	}
	if err := tc.parseExpected(&outLines[0]); err != nil {
		return err
	}
	return nil
}

func readFile(filePath string) ([]string, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	res := strings.Split(string(f), "\r\n")
	return res, nil
}
