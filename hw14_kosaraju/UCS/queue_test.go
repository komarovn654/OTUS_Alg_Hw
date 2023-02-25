package ucs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
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
			array:  []int{1, 2, 3, 4, 5, 6, 7},
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
			queue := NewQueue()

			for _, v := range tc.array {
				queue.Queue(v)
			}

			for _, v := range tc.expect {
				value, _ := queue.Dequeue()
				require.Equal(t, v, value)
			}
		})
	}
}
