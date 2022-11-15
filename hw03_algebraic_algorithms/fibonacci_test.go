package hw03algebraicalgorithms

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	n      int
	expect *big.Int
}

func readLine(filePath string) (string, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return strings.Split(string(f), "\n")[0], nil
}

func prepareTest(testNum int) (testStruct, error) {
	in := fmt.Sprintf("Fibo//test.%v.in", testNum)
	out := fmt.Sprintf("Fibo//test.%v.out", testNum)

	i, _ := readLine(in)
	o, _ := readLine(out)

	n, _ := strconv.Atoi(i)

	bio := big.NewInt(0)
	expect, _ := bio.SetString(o, 10)

	return testStruct{n, expect}, nil
}

func TestFib(t *testing.T) {
	// for _, tc := range tests {
	// 	t.Run(tc.name+" Recursive", func(t *testing.T) {
	// 		require.Equal(t, tc.expected, FibRecursive(tc.in), "Recursive")
	// 	})
	// 	t.Run(tc.name+" Iterative", func(t *testing.T) {
	// 		require.Equal(t, tc.expected, FibIterative(tc.in), "Iterative")
	// 	})
	// 	t.Run(tc.name+" Golden ratio", func(t *testing.T) {
	// 		require.Equal(t, tc.expected, FibGolden(tc.in), "Golden ratio")
	// 	})
	// }
	for i := 0; i < 13; i++ {
		ts, err := prepareTest(i)
		if err != nil {
			log.Fatal(err)
		}

		name := fmt.Sprintf("N=%v", ts.n)
		// t.Run(name+" Recursive", func(t *testing.T) {
		// 	require.Equal(t, ts.expect, FibRecursive(ts.n), "Recursive")
		// })
		t.Run(name+" Iterative", func(t *testing.T) {
			require.Equal(t, ts.expect, FibIterative(ts.n), "Iterative")
		})
		// t.Run(name+" Golden ratio", func(t *testing.T) {
		// 	require.Equal(t, ts.expect, FibGolden(ts.n), "Golden ratio")
		// })
		// t.Run(name+" Matrix", func(t *testing.T) {
		// 	require.Equal(t, ts.expect, FibMatrix(ts.n), "Matrix")
		// })
	}
}
