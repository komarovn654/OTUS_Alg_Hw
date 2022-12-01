package hw04arrays

import (
	"fmt"
)

type PriorityQueue interface {
	Enqueue(priority int, t Item)
	Dequeue() Item
}

type pqArray struct {
	lists FactorArray
}

func (pq *pqArray) Enqueue(priority int, t Item) {
	if priority > pq.lists.Size() {
		pq.lists.resize(priority + 1)
	}

	if pq.lists.Get(priority) == nil {
		pq.lists.Set(NewList(), priority)
	}

	pq.pushList(priority, t)
}

func (pq *pqArray) Dequeue() Item {
	if pq.lists.Size() == 0 {
		return nil
	}

	i := 0
	for i = 0; i < pq.lists.Size() && pq.isListEmpty(i); i++ {
	}

	return pq.popList(i)
}

func (pq *pqArray) isListEmpty(listNum int) bool {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		return l.Len() <= 0
	}

	return true
}

func (pq *pqArray) popList(listNum int) Item {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		v := l.Front()
		l.Remove(v)
		return v.Value
	}
	return nil
}

func (pq *pqArray) pushList(listNum int, t Item) {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		l.PushBack(t)
		return
	}
}

func NewPQArray() pqArray {
	return pqArray{}
}

type pqList struct {
	lists List
}

type store struct {
	v    Item
	prio int
}

func (pq *pqList) Enqueue(priority int, t Item) {
	if pq.lists.Front() == nil {
		pq.pushToNewList(store{prio: priority, v: t}, 0)
		return
	}

	list := pq.lists.Front()
	for i := 0; i < pq.lists.Len(); i++ {
		prio := getListPriority(list)
		if prio > priority {
			pq.pushToNewList(store{prio: priority, v: t}, i)
			return
		} else if prio == priority {
			pq.pushToExistList(store{prio: priority, v: t}, list)
			return
		}
		list = list.Next
	}

	pq.pushToNewList(store{prio: priority, v: t}, pq.lists.Len())
}

func (pq *pqList) Dequeue() Item {
	if pq.lists.Front() == nil { // empty list
		return nil
	}

	frontList := pq.lists.Front()
	l, ok := frontList.Value.(*list)
	if !ok {
		panic(fmt.Errorf("cannot cast to list"))
	}

	storeItem := l.Back()
	v, ok := storeItem.Value.(store)
	if !ok {
		panic(fmt.Errorf("cannot cast to store item"))
	}

	l.Remove(storeItem)
	if l.len == 0 {
		pq.lists.Remove(frontList)
	}

	return v.v
}

func getListPriority(l *ListItem) int {
	list, ok := l.Value.(*list)
	if !ok {
		panic(fmt.Errorf("cannot cast to list"))
	}
	if storeItem, ok := list.Front().Value.(store); ok {
		return storeItem.prio

	}
	panic(fmt.Errorf("cannot cast to store item"))
}

func (pq *pqList) pushToNewList(item store, listPos int) {
	list := NewList()
	list.PushFront(item)
	pq.lists.Insert(listPos, list)
}

func (pq *pqList) pushToExistList(item store, l *ListItem) {
	list, ok := l.Value.(*list)
	if !ok {
		panic(fmt.Errorf("cannot cast to list"))
	}
	list.PushFront(item)
}

func (pq *pqList) PrintPQList() {
	mylist := pq.lists.Front()
	for i := 0; i < pq.lists.Len(); i++ {
		l, _ := mylist.Value.(*list)
		item := l.Front()
		for j := 0; j < l.Len(); j++ {
			v, _ := item.Value.(store)
			fmt.Printf("value: %v; priority: %v\n", v.v, v.prio)
			item = item.Next
		}
		mylist = mylist.Next
	}
}

func NewPQList() pqList {
	return pqList{lists: NewList()}
}
