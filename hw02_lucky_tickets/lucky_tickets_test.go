package hw02_lucky_tickets

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	n       int
	tickets int
}

func prepareTest(testNum int) (testStruct, error) {
	in := fmt.Sprintf("tickets//test.%v.in", testNum)
	out := fmt.Sprintf("tickets//test.%v.out", testNum)
	inFile, err := os.ReadFile(in)
	if err != nil {
		return testStruct{}, err
	}
	outFile, err := os.ReadFile(out)
	if err != nil {
		return testStruct{}, err
	}
	outLines := strings.Split(string(outFile), "\n")

	n, err := strconv.Atoi(string(inFile))
	if err != nil {
		return testStruct{}, err
	}
	expectedCount, err := strconv.Atoi(string(outLines[0]))
	if err != nil {
		return testStruct{}, err
	}
	return testStruct{n, expectedCount}, nil
}

func TestLuckyTickets(t *testing.T) {
	for i := 0; i < 10; i++ {
		ts, err := prepareTest(i)
		if err != nil {
			log.Fatal(err)
		}
		t.Run(fmt.Sprintf("N/2=%v", ts.n), func(t *testing.T) {
			require.Equal(t, ts.tickets, LuckyTickets(ts.n))
		})
	}
}

func TestCalculateSumWrap(t *testing.T) {
	tests := []struct {
		name    string
		refSums []int
		n       int
		sums    []int
		sum     int
	}{
		{
			name:    "N=2",
			refSums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			n:       2,
			sums:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			sum:     670,
		},
		{
			name: "N=6",
			refSums: []int{1, 5, 15, 35, 70, 126, 210, 330, 495, 715, 996, 1340, 1745, 2205, 2710, 3246, 3795, 4335,
				4840, 5280, 5631, 5875, 6000, 6000, 5875, 5631, 5280, 4840, 4335, 3795, 3246, 2710, 2205, 1745, 1340,
				996, 715, 495, 330, 210, 126, 70, 35, 15, 5, 1},
			n: 6,
			sums: []int{1, 6, 21, 56, 126, 252, 462, 792, 1287, 2002, 2997, 4332, 6062, 8232, 10872, 13992, 17577, 21582,
				25927, 30492, 35127, 39662, 43917, 47712, 50877, 53262, 54747, 55252, 54747, 53262, 50877, 47712, 43917,
				39662, 35127, 30492, 25927, 21582, 17577, 13992, 10872, 8232, 6062, 4332, 2997, 2002, 1287, 792, 462, 252,
				126, 56, 21, 6, 1},
			sum: 39581170420,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sums, sum := calculateSumWrap(tc.refSums, tc.n)
			require.Equal(t, sums, tc.sums)
			require.Equal(t, sum, tc.sum)
		})
	}
}
