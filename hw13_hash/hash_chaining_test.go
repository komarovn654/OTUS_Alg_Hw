package hw13hash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func nodeFromKeys(keys []key) node {

	for i, k := range keys {

	}
}

func TestGetHashCode(t *testing.T) {
	tests := []struct {
		name string
		k    key
		code int
	}{
		{
			name: "common case",
			k:    "qwerty",
			code: 684,
		},
		{
			name: "empty key",
			k:    "",
			code: 0,
		},
		{
			name: "single char",
			k:    "a",
			code: 97,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.code, tc.k.getHashCode())
		})
	}
}

func TestIsKeyExists(t *testing.T) {
	tests := []struct {
		name     string
		nodeKeys []key
		k        key
	}{
		{
			name:     "common case",
			nodeKeys: []key{"qwerty", "tmp", "", "foo", "hash"},
			k:        "foo",
		},
	}
	fmt.Println(nodeFromKeys(tests[0].nodeKeys))
	// for _, tc := range tests {
	// 	t.Run(tc.name, func(t *testing.T) {
	// 		require.Equal(t, tc.code, tc.k.getHashCode())
	// 	})
	// }
}
