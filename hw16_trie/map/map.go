package trie_map

import (
	trie "github.com/komarovn654/OTUS_Alg_Hw/hw16_trie"
)

type trieMap struct {
	trie trie.Trie
}

type Item struct {
}

func Init() trieMap {
	return trieMap{}
}

func (m *trieMap) Add(item Item) {

}

func (m *trieMap) Get() (exist bool, value any) {
	return false, nil
}
