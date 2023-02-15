package hw14_kosaraju

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	test := []int{1, 2, 3, 4, 5, 6, 7}
	stack := Stack()

	for _, v := range test {
		stack.Push(v)
	}

	f := func(item any) {
		fmt.Println(*(&item))
	}

	stack.forEach(f)

	for range test {
		fmt.Println(stack.Pop())
	}
}
