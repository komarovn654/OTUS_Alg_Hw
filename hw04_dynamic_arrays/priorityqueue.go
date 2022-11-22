package hw04arrays

type PriorityQueue struct {
	lists FactorArray
}

func (pq *PriorityQueue) Enqueue(priority int, t Item) {
	if priority > pq.lists.Size() {
		pq.lists.resize(priority)

	}

	if pq.lists.Get(priority) == nil {
		pq.lists.Set(NewList(), priority)
	}

	pq.pushList(priority, t)
}

func (pq *PriorityQueue) Dequeue() Item {
	if pq.lists.Size() == 0 {
		return nil
	}

	i := 0
	for i = 0; i < pq.lists.Size() && pq.isListEmpty(i); i++ {
	}

	return pq.popList(i)
}

func (pq *PriorityQueue) isListEmpty(listNum int) bool {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		return l.Len() <= 0
	}

	return true
}

func (pq *PriorityQueue) popList(listNum int) Item {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		v := l.Front()
		l.Remove(v)
		return v.Value
	}
	return nil
}

func (pq *PriorityQueue) pushList(listNum int, t Item) {
	pl := pq.lists.Get(listNum)
	if l, ok := pl.(List); ok {
		l.PushBack(t)
		return
	}
}
