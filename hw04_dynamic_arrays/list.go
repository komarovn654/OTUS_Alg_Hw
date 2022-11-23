package hw04arrays

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	Insert(pos int, i interface{})
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	frontItem *ListItem
	backItem  *ListItem
	len       int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.frontItem
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.backItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  l.frontItem,
		Prev:  nil,
	}

	l.len++
	l.frontItem.Prev = newItem
	l.frontItem = newItem
	if l.len == 1 {
		newItem.Next = nil
		l.backItem = newItem
	}
	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.backItem,
	}

	l.len++
	l.backItem.Next = newItem
	l.backItem = newItem
	if l.len == 1 {
		newItem.Prev = nil
		l.frontItem = newItem
	}
	return newItem
}

func (l *list) Remove(i *ListItem) {
	l.len--
	if l.len == 0 {
		l.frontItem = new(ListItem)
		l.backItem = new(ListItem)
		return
	}

	if i.Next == nil {
		i.Prev.Next = nil
		l.backItem = i.Prev
		return
	}

	if i.Prev == nil {
		i.Next.Prev = nil
		l.frontItem = i.Next
		return
	}

	i.Next.Prev = i.Prev
	i.Prev.Next = i.Next
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func (l *list) Insert(pos int, i interface{}) {
	item := l.Front()
	for i := 0; i < pos-1; i++ {
		item = item.Next
	}

	newItem := &ListItem{
		Value: i,
		Next:  item.Next,
		Prev:  item,
	}

	item.Next.Prev = newItem
	item.Next = newItem
	l.len++
}

func NewList() List {
	return &list{
		frontItem: new(ListItem),
		backItem:  new(ListItem),
	}
}
