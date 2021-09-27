package stack

import "errors"

type LinkedListStack struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func NewLinkedListStack() *LinkedListStack {
	s := new(LinkedListStack)
	s.head = nil
	return s
}

func (s *LinkedListStack) Push(value int) {
	node := new(Node)
	node.value = value
	node.next = s.head
	s.head = node
}

func (s *LinkedListStack) Pop() (int, error) {
	if s.head == nil {
		return 0, errors.New("Stack empty")
	}
	node := s.head
	s.head = s.head.next

	defer func() {
		node = nil
	}()
	return node.value, nil
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.head == nil
}

func (s *LinkedListStack) Size() int {
	count := 0
	head := s.head
	for head != nil {
		head = head.next
		count++
	}
	return count
}
