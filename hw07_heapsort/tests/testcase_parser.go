package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/komarovn654/OTUS_Alg_Hw/hw07_heapsort"
)

type testCase struct {
	Name     string
	N        int64
	Array    []hw07_heapsort.Item
	Expected []hw07_heapsort.Item
}

func (tc *testCase) parseN(str *string) (err error) {
	tc.N, err = strconv.ParseInt(*str, 10, 64)
	return err
}

func (tc *testCase) parseArray(str *string) (err error) {
	tc.Array = make([]hw07_heapsort.Item, tc.N)
	for i, item := range strings.Split(*str, " ") {
		value, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return err
		}
		tc.Array[i] = hw07_heapsort.Item(value)
	}
	return nil
}

func (tc *testCase) parseExpected(str *string) (err error) {
	tc.Expected = make([]hw07_heapsort.Item, tc.N)
	for i, item := range strings.Split(*str, " ") {
		value, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return err
		}
		tc.Expected[i] = hw07_heapsort.Item(value)
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
	// f, err := os.Open(filePath)
	// if err != nil {
	// 	return nil, err
	// }
	// defer f.Close()

	// str := make([]string, 0)
	// fscan := bufio.NewScanner(f)
	// for fscan.Scan() {
	// 	str = append(str, fscan.Text())
	// }
	// fmt.Println(filePath, len(str))
	// return str, nil
	f, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	res := strings.Split(string(f), "\n")
	return res, nil
}
