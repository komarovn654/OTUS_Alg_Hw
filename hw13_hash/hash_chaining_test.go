package hw13hash

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

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

func generateRandomItems(len int) []tableItem {
	ar := make([]tableItem, len)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range ar {
		k := key(strconv.Itoa(r1.Intn(10_000)))
		v := value(r1.Intn(10_000))

		ar[i] = tableItem{k, v}
	}

	return ar
}

func TestChainingHashTable(t *testing.T) {
	t.Run("", func(t *testing.T) {
		items := generateRandomItems(1000)
		checked := make(map[key]bool)
		ht := InitChainigHashTable(10)

		for _, item := range items {
			ht.Set(item)
		}

		for i := len(items) - 1; i > 0; i-- {
			if checked[items[i].k] {
				// Item был перезаписан
				continue
			}

			v, ok := ht.Get(items[i].k)
			require.True(t, ok)
			if items[i].v != v {
				for i, item := range items {
					fmt.Println(i, item, v)
				}
			}
			require.Equal(t, items[i].v, v)
			checked[items[i].k] = true
		}

		for _, item := range items {
			ht.Remove(item.k)
			v, ok := ht.Get(item.k)
			require.False(t, ok)
			require.Equal(t, value(0), v)
		}
	})
}

func TestRemove(t *testing.T) {
	tests := []struct {
		name  string
		items []tableItem
		exist bool
	}{
		{
			name:  "no collisions",
			items: []tableItem{{"a", 46}, {"b", 68}, {"c", 9}, {"d", 2}, {"e", 5}},
			exist: true,
		},
		{
			name:  "collisions",
			items: []tableItem{{"abc", 46}, {"b", 68}, {"bac", 9}, {"cab", 2}},
			exist: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ht := InitChainigHashTable(10)
			for _, item := range tc.items {
				ht.Set(item)
			}

			for _, item := range tc.items {
				ht.Remove(item.k)
				_, ok := ht.Get(item.k)
				require.False(t, ok)
			}
		})
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		name  string
		items []tableItem
		exist bool
	}{
		{
			name:  "no collisions",
			items: []tableItem{{"a", 46}, {"b", 68}, {"c", 9}, {"d", 2}, {"e", 5}},
			exist: true,
		},
		{
			name:  "collisions",
			items: []tableItem{{"abc", 46}, {"b", 68}, {"bac", 9}, {"cab", 2}},
			exist: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ht := InitChainigHashTable(10)
			for _, item := range tc.items {
				ht.Set(item)
			}

			for _, item := range tc.items {
				value, ok := ht.Get(item.k)
				require.True(t, ok)
				require.Equal(t, item.v, value)
			}
		})
	}

	t.Run("doesnt exist key", func(t *testing.T) {
		ht := InitChainigHashTable(10)
		for _, item := range tests[0].items {
			ht.Set(item)
		}

		val, ok := ht.Get("lkj")
		require.False(t, ok)
		require.Equal(t, value(0), val)
	})
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
			tableSize: 10,
		},
		{
			name:      "collisions",
			items:     []tableItem{{"abc", 46}, {"b", 68}, {"bac", 9}, {"cab", 2}},
			tableSize: 10,
		},
		{
			name:      "same keys",
			items:     []tableItem{{"abc", 46}, {"abc", 68}, {"abc", 9}, {"abc", 2}},
			tableSize: 10,
		},
	}

	// testcases without rehash
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ht := InitChainigHashTable(tc.tableSize)
			for _, item := range tc.items {
				index := item.k.getHashCode() % ht.size

				ht.Set(item)

				require.Equal(t, ht.table[index].item, item)
				require.Equal(t, ht.size, tc.tableSize)
			}
		})
	}

	// rehash
	t.Run("rehash", func(t *testing.T) {
		initSize := 2
		expectSize := 5 // 2n + 1

		ht := InitChainigHashTable(initSize)
		for i := 0; i < rehashLoadFactor+1; i++ {
			ht.Set(tableItem{key(strconv.Itoa(i)), value(i)})
			ht.Set(tableItem{key(strconv.Itoa(i + 10)), value(i + 10)})
		}

		require.Equal(t, expectSize, ht.size)

		for i := 0; i < rehashLoadFactor+1; i++ {
			item := tableItem{key(strconv.Itoa(i)), value(i)}
			index := item.k.getHashCode() % ht.size

			i, ok := ht.table[index].isKeyExist(item.k)
			require.True(t, ok)
			require.Equal(t, item, i)
		}
	})
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

			item, ok := node.isKeyExist(tc.k)
			require.Equal(t, tc.res, ok)
			if ok {
				require.Equal(t, tableItem{tc.k, 0}, item)
			}
		})
	}
}
