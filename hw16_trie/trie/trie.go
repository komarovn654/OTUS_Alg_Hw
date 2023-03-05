package trie

type Trie struct {
	prefix [128]*layer
}

type layer struct {
	prefix  [128]*layer
	isExist bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	l := &this.prefix

	for i, char := range word {
		if l[char] == nil {
			l[char] = new(layer)
		}

		if i == len(word)-1 {
			l[char].isExist = true
			return
		}

		l = &l[char].prefix
	}
}

func (this *Trie) Search(word string) bool {
	l := this.prefix

	for i, char := range word {
		if l[char] == nil {
			return false
		}

		if i == len(word)-1 {
			return l[char].isExist
		}

		l = l[char].prefix
	}

	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	l := this.prefix

	for i, char := range prefix {
		if l[char] == nil {
			return false
		}

		if i == len(prefix)-1 {
			return true
		}

		l = l[char].prefix
	}

	return false
}
