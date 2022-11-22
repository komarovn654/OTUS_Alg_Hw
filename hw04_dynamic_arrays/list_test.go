package hw04arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()
		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func checkIntList(l List, expectedValues []int) bool {
	item := l.Front()
	i := 0
	for ; item != nil; i++ {
		if item.Value.(int) != expectedValues[i] {
			return false
		}
		item = item.Next
	}
	return true
}

func TestLen(t *testing.T) {
	l := NewList()
	require.Equal(t, 0, l.Len())

	l.PushFront(10)
	require.Equal(t, 1, l.Len())
	l.PushBack(20)
	require.Equal(t, 2, l.Len())

	l.Remove(l.Back())
	require.Equal(t, 1, l.Len())

	l.Remove(l.Back())
	require.Equal(t, 0, l.Len())

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			l.PushFront(i)
		} else {
			l.PushBack(i)
		}
	}
	require.Equal(t, 10, l.Len())

	l.Back()
	require.Equal(t, 10, l.Len())
	l.Front()
	require.Equal(t, 10, l.Len())
	l.MoveToFront(l.Back())
	require.Equal(t, 10, l.Len())
}

func TestFront(t *testing.T) {
	l := NewList()
	require.Nil(t, l.Front())

	testItem := l.PushBack(10)
	require.Equal(t, testItem, l.Front())
	l.Remove(l.Back())

	testItem = l.PushFront(20)
	require.Equal(t, testItem, l.Front())
	l.PushBack(30)
	require.Equal(t, testItem, l.Front())

	testItem = l.PushFront(40)
	require.Equal(t, testItem, l.Front())
}

func TestBack(t *testing.T) {
	l := NewList()
	require.Nil(t, l.Back())

	testItem := l.PushFront(10)
	require.Equal(t, testItem, l.Back())
	l.Remove(l.Front())

	testItem = l.PushBack(20)
	require.Equal(t, testItem, l.Back())
	l.PushFront(30)
	require.Equal(t, testItem, l.Back())

	testItem = l.PushBack(40)
	require.Equal(t, testItem, l.Back())
}

func TestPushFront(t *testing.T) {
	l := NewList()
	testItem := l.PushFront(10)
	require.Equal(t, testItem, l.Front())
	require.True(t, checkIntList(l, []int{10}))

	for i := 0; i < 5; i++ {
		testItem = l.PushFront(i)
		require.Equal(t, testItem, l.Front())
	}
	require.True(t, checkIntList(l, []int{4, 3, 2, 1, 0, 10}))
}

func TestPushBack(t *testing.T) {
	l := NewList()
	testItem := l.PushBack(10)
	require.Equal(t, testItem, l.Back())
	require.True(t, checkIntList(l, []int{10}))

	for i := 0; i < 5; i++ {
		testItem := l.PushBack(i)
		require.Equal(t, testItem, l.Back())
	}
	require.True(t, checkIntList(l, []int{10, 0, 1, 2, 3, 4}))
}

func TestMoveToFront(t *testing.T) {
	l := NewList()

	frontItem := l.PushFront(10)
	l.MoveToFront(frontItem)
	require.Equal(t, frontItem.Value, l.Front().Value)
	require.True(t, checkIntList(l, []int{10}))

	backItem := l.PushBack(20)
	l.MoveToFront(l.Front())
	require.Equal(t, frontItem.Value, l.Front().Value)
	require.True(t, checkIntList(l, []int{10, 20}))
	l.MoveToFront(l.Back())
	require.Equal(t, backItem.Value, l.Front().Value)
	require.True(t, checkIntList(l, []int{20, 10}))

	var testItem *ListItem
	for i := 0; i < 5; i++ {
		if i == 3 {
			testItem = l.PushBack(i)
			continue
		}
		l.PushBack(i)
	}
	l.MoveToFront(testItem)
	require.Equal(t, testItem.Value, l.Front().Value)
	require.True(t, checkIntList(l, []int{3, 20, 10, 0, 1, 2, 4}))

	l.MoveToFront(l.Front())
	require.True(t, checkIntList(l, []int{3, 20, 10, 0, 1, 2, 4}))
	l.MoveToFront(l.Back())
	require.True(t, checkIntList(l, []int{4, 3, 20, 10, 0, 1, 2}))
}
