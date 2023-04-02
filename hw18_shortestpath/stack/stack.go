package stack

type Stack struct {
	head *node
}

type node struct {
	item any
	next *node
}

func New() Stack {
	return Stack{}
}

func (s *Stack) Push(new any) {
	s.head = &node{item: new, next: s.head}
}

func (s *Stack) Pop() (any, bool) {
	if s.head == nil {
		return nil, false
	}

	tmp := s.head.item
	s.head = s.head.next
	return tmp, true
}

func (s *Stack) ForEach(f func(any)) {
	tmp := s.head

	for ; tmp != nil; tmp = tmp.next {
		f(tmp.item)
	}
}
