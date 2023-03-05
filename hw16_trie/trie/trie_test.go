package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type result any

func TestTire(t *testing.T) {
	tests := []struct {
		commands []string // "Insert", "Search", "StartsWith", "Construct"
		args     []string
		expect   []result
	}{
		{
			commands: []string{"Construct", "Insert", "Search", "Search", "StartsWith", "Insert", "Search"},
			args:     []string{"", "apple", "apple", "app", "app", "app", "app"},
			expect:   []result{nil, nil, true, false, true, nil, true},
		},
		{
			commands: []string{"Construct", "Insert", "StartsWith"},
			args:     []string{"", "a", "aa"},
			expect:   []result{nil, nil, false},
		},
	}

	for _, tc := range tests {
		t.Run("common case", func(t *testing.T) {
			var trie Trie
			for i, command := range tc.commands {
				switch command {
				case "Insert":
					trie.Insert(tc.args[i])
				case "Search":
					res := trie.Search(tc.args[i])
					require.Equal(t, tc.expect[i], res)
				case "StartsWith":
					res := trie.StartsWith(tc.args[i])
					require.Equal(t, tc.expect[i], res)
				case "Construct":
					trie = Constructor()
				}
			}
		})
	}
}
