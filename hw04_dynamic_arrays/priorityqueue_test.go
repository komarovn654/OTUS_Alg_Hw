package hw04arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type val struct {
	prio  int
	value interface{}
}

func TestEnqueueArray(t *testing.T) {
	tests := []struct {
		name       string
		pushValues []val
		expect     []int
	}{
		{
			name:       "single priority",
			pushValues: []val{{10, 423}, {10, 32}, {10, 5}, {10, 12}, {10, 3}, {10, 42}, {10, 523}, {10, 523}, {10, 6}},
			expect:     []int{423, 32, 5, 12, 3, 42, 523, 523, 6},
		},
		{
			name:       "common case",
			pushValues: []val{{10, 423}, {1, 32}, {3, 5}, {4, 12}, {1, 3}, {2, 42}, {3, 523}, {10, 523}, {1, 6}, {11, 35}},
			expect:     []int{32, 3, 6, 42, 5, 523, 12, 423, 523, 35},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+" PQArray", func(t *testing.T) {
			pq := NewPQArray()
			for _, v := range tc.pushValues {
				pq.Enqueue(v.prio, v.value)
			}
			for i := range tc.pushValues {
				require.Equal(t, tc.expect[i], pq.Dequeue())
			}
		})

		t.Run(tc.name+" PQList", func(t *testing.T) {
			pq := NewPQList()
			for _, v := range tc.pushValues {
				pq.Enqueue(v.prio, v.value)
			}
			//pq.PrintPQList()
			for i := range tc.pushValues {
				require.Equal(t, tc.expect[i], pq.Dequeue())
			}
		})
	}
}
