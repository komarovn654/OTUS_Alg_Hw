package trie_map

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrieMap(t *testing.T) {
	tests := []struct {
		items  []Item
		search []string
		expect []struct {
			exist bool
			value any
		}
	}{
		{items: []Item{{"app", 12}, {"app", 13}, {"apple", 16}, {"key", 1}, {"car", 51}, {"kevin", 5}, {"apple", 14}},
			search: []string{"app", "apple", "key", "car", "kevin", "cap"},
			expect: []struct {
				exist bool
				value any
			}{
				{true, 13}, {true, 14}, {true, 1}, {true, 51}, {true, 5}, {false, nil},
			},
		},
	}

	for _, tc := range tests {
		t.Run("common", func(t *testing.T) {
			m := Constructor()

			for _, item := range tc.items {
				m.Insert(item)
			}

			for i, key := range tc.search {
				value, exist := m.Search(key)
				require.Equal(t, tc.expect[i].exist, exist)
				require.Equal(t, tc.expect[i].value, value)
			}

			for _, key := range tc.search {
				m.Remove(key)
				value, exist := m.Search(key)
				require.False(t, exist)
				require.Nil(t, value)
			}
		})
	}
}
