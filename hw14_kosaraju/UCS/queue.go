package ucs

type List struct {
	frontItem *node
	backItem  *node
}

type node struct {
	item any
	next *node
	prev *node
}

func NewQueue() List {
	return List{}
}

func (l *List) Queue(item any) {
	newItem := &node{
		item: item,
		next: nil,
		prev: l.backItem,
	}

	if l.backItem == nil || l.frontItem == nil {
		l.frontItem = newItem
		l.backItem = newItem
		return
	}

	l.backItem.next = newItem
	l.backItem = newItem
}

func (l *List) Dequeue() (any, bool) {
	if l.frontItem == nil || l.backItem == nil {
		return nil, false
	}

	tmp := l.frontItem.item

	l.frontItem = l.frontItem.next

	return tmp, true
}
