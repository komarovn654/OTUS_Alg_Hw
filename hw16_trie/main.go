package main

import "fmt"

const (
	EMPTY = iota
	PREFIX
	WORD
)

type Trie struct {
	prefix [128]*layer
}

type layer struct {
	prefix [128]*layer
	exist  uint8
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

		if l[char].exist == EMPTY {
			l[char].exist = PREFIX
		}

		if i == len(word)-1 {
			l[char].exist = WORD
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
			return l[char].exist == WORD
		}

		if l[char].exist == EMPTY {
			return false
		}

		l = l[char].prefix
	}

	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	var char rune
	l := this.prefix

	for _, char = range prefix {
		if l[char] == nil {
			return false
		}
		if l[char].exist == EMPTY {
			return false
		}
		l = l[char].prefix
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	t := Constructor()
	t.Insert("apple")
	fmt.Println(t.Search("apple"))
}
