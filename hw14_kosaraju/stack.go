package hw14_kosaraju

type stack struct {
	head *node
}

type node struct {
	item any
	next *node
}

func Stack() stack {
	return stack{}
}

func (s *stack) Push(new any) {
	s.head = &node{item: new, next: s.head}
}

func (s *stack) Pop() any {
	if s.head == nil {
		return nil
	}

	tmp := s.head.item
	s.head = s.head.next
	return tmp
}

func (s *stack) forEach(f func(any)) {
	tmp := s.head

	for ; tmp != nil; tmp = tmp.next {
		f(tmp.item)
	}
}
