package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	in     uint64 // input value
	expCnt uint64 // expected count
	expPos uint64 // expected positions
}

func filesCount(dir string) (count int, err error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return count, err
	}

	for _, f := range files {
		if strings.Contains(f.Name(), "test") {
			count++
		}
	}

	return count, err
}

func readFile(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	str := make([]string, 0)
	fscan := bufio.NewScanner(f)
	for fscan.Scan() {
		str = append(str, fscan.Text())
	}

	return str, nil
}

func parseKingTC(dir string) ([]testCase, error) {
	tcNum, err := filesCount(dir)
	os.Chdir(dir)
	if err != nil {
		return nil, err
	}
	testCases := make([]testCase, tcNum/2)

	for i, tc := range testCases {
		inLines, err := readFile(fmt.Sprintf("test.%v.in", i))
		if err != nil {
			return nil, err
		}
		outLines, err := readFile(fmt.Sprintf("test.%v.out", i))
		if err != nil {
			return nil, err
		}

		tc.in, err = strconv.ParseUint(inLines[0], 10, 64)
		if err != nil {
			return nil, err
		}
		tc.expCnt, err = strconv.ParseUint(outLines[0], 10, 64)
		if err != nil {
			return nil, err
		}
		tc.expPos, err = strconv.ParseUint(outLines[1], 10, 64)
		if err != nil {
			return nil, err
		}
		testCases[i] = tc
	}
	return testCases, nil
}
