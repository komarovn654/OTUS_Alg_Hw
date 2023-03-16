package trie_map

type trieMap struct {
	prefix [128]*layer
}

type layer struct {
	prefix [128]*layer
	value  any
}

type Item struct {
	Key   string
	Value any
}

func Constructor() trieMap {
	return trieMap{}
}

func (this *trieMap) Insert(item Item) {
	l := &this.prefix

	for i, char := range item.Key {
		if l[char] == nil {
			l[char] = new(layer)
		}

		if i == len(item.Key)-1 {
			l[char].value = item.Value
			return
		}

		l = &l[char].prefix
	}
}

func (this *trieMap) Search(key string) (value any, exist bool) {
	l := this.prefix

	for i, char := range key {
		if l[char] == nil {
			return nil, false
		}

		if i == len(key)-1 {
			return l[char].value, l[char].value != nil
		}

		l = l[char].prefix
	}

	return nil, false
}

func (this *trieMap) Remove(key string) {
	l := this.prefix

	for i, char := range key {
		if l[char] == nil {
			return
		}

		if i == len(key)-1 {
			l[char].value = nil
			return
		}

		l = l[char].prefix
	}

	return
}
