package hw04arrays

import (
	"fmt"
)

type PriorityQueue interface {
	Enqueue(priority int, t Item)
	Dequeue() Item
}

type PQArray struct {
	lists FactorArray
}

func (pq *PQArray) Enqueue(priority int, t Item) {
	if priority > pq.lists.Size() {
		pq.lists.resize(priority)

	}

	if pq.lists.Get(priority) == nil {
		pq.lists.Set(NewList(), priority)
	}

	pq.pushList(priority, t)
}

func (pq *PQArray) Dequeue() Item {
	if pq.lists.Size() == 0 {
		return nil
	}

	i := 0
	for i = 0; i < pq.lists.Size() && pq.isListEmpty(i); i++ {
	}

	return pq.popList(i)
}

func (pq *PQArray) isListEmpty(listNum int) bool {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		return l.Len() <= 0
	}

	return true
}

func (pq *PQArray) popList(listNum int) Item {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		v := l.Front()
		l.Remove(v)
		return v.Value
	}
	return nil
}

func (pq *PQArray) pushList(listNum int, t Item) {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		l.PushBack(t)
		return
	}
}

type PQList struct {
	lists list
}

type store struct {
	v    Item
	prio int
}

func (pq *PQList) Enqueue(priority int, t Item) {
	if pq.lists.len == 0 {
		newList := NewList()
		newList.PushFront(store{v: t, prio: priority})
		pq.lists.PushFront(newList)
		return
	}

	list := pq.lists.Front()
	for i := 0; i < pq.lists.Len(); i++ {
		item := pq.getFrontItem(list)
		if item.prio > priority {
			newList := NewList()
			newList.PushFront(store{v: t, prio: priority})
			pq.lists.Insert(i, newList)
		}
	}

}

func (pq *PQList) Dequeue() Item {
	return nil
}

func (pq *PQList) getFrontItem(l *ListItem) store {
	list, ok := l.Value.(list)
	if !ok {
		panic(fmt.Errorf("cannot cast to list"))
	}
	if storeItem, ok := list.Front().Value.(store); ok {
		return storeItem

	}
	panic(fmt.Errorf("cannot cast to store item"))
}
