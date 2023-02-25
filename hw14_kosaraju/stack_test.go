package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPushPop(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		expect []int
	}{
		{
			name:   "empty",
			array:  nil,
			expect: nil,
		},
		{
			name:   "common",
			array:  []int{1, 2, 3, 4, 5, 6, 7}, // reverse when push
			expect: []int{7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:   "single node",
			array:  []int{1},
			expect: []int{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stack := New()

			for _, v := range tc.array {
				stack.Push(v)
			}

			for _, v := range tc.expect {
				value, _ := stack.Pop()
				require.Equal(t, v, value)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		expect []int
	}{
		{
			name:   "empty",
			array:  nil,
			expect: nil,
		},
		{
			name:   "common",
			array:  []int{1, 2, 3, 4, 5, 6, 7}, // reverse when push
			expect: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:   "single node",
			array:  []int{1},
			expect: []int{1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			stack := New()
			for _, v := range tc.array {
				stack.Push(v)
			}
			stack.Reverse()

			for _, v := range tc.expect {
				value, _ := stack.Pop()
				require.Equal(t, v, value)
			}
		})
	}
}
