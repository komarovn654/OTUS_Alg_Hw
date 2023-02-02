package hw13hash

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func nodeFromKeys(keys []key) node {
	if len(keys) == 0 {
		return node{}
	}

	head := node{item: tableItem{k: keys[0]}, next: nil}
	for _, k := range keys[1:] {
		newHead := node{item: tableItem{k: k}, next: &node{item: head.item, next: head.next}}
		head = newHead
	}

	return head
}

func TestSet(t *testing.T) {
	tests := []struct {
		name      string
		items     []tableItem
		tableSize int
	}{
		{
			name:      "no collisions",
			items:     []tableItem{{"a", 46}, {"b", 68}, {"c", 9}, {"d", 2}, {"e", 5}},
			tableSize: initialTableSize,
		},
		{
			name:      "collisions",
			items:     []tableItem{{"abc", 46}, {"b", 68}, {"bac", 9}, {"cab", 2}},
			tableSize: initialTableSize,
		},
		{
			name:      "same keys",
			items:     []tableItem{{"abc", 46}, {"abc", 68}, {"abc", 9}, {"abc", 2}},
			tableSize: initialTableSize,
		},
		// {
		// 	name:      "rehash",
		// 	items:     []tableItem{},
		// 	tableSize: ,
		// },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ht := initChainigHashTable()
			for _, item := range tc.items {
				index := item.k.getHashCode() % ht.size

				ht.Set(item)

				require.Equal(t, ht.table[index].item, item)
				require.Equal(t, ht.size, tc.tableSize)
			}
		})
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
		res      bool
	}{
		{
			name:     "common case exists",
			nodeKeys: []key{"qwerty", "tmp", "", "foo", "hash"},
			k:        "foo",
			res:      true,
		},
		{
			name:     "common case exists, last",
			nodeKeys: []key{"qwerty", "tmp", "", "hash", "foo"},
			k:        "foo",
			res:      true,
		},
		{
			name:     "common case exists, first",
			nodeKeys: []key{"foo", "qwerty", "tmp", "", "hash"},
			k:        "foo",
			res:      true,
		},
		{
			name:     "common case doesnt exists",
			nodeKeys: []key{"qwerty", "tmp", "", "hash"},
			k:        "foo",
			res:      false,
		},
		{
			name:     "empty node",
			nodeKeys: []key{},
			k:        "foo",
			res:      false,
		},
		{
			name:     "single child node",
			nodeKeys: []key{"foo"},
			k:        "foo",
			res:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			node := nodeFromKeys(tc.nodeKeys)

			require.Equal(t, tc.res, node.isKeyExist(tc.k))
		})
	}
}
