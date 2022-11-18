package _8_stack

import "fmt"

/*
基于链表实现的栈
*/
type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	// 栈顶节点
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.topNode == nil
}

func (s *LinkedListStack) Push(v interface{}) {
	s.topNode = &node{next: s.topNode, val: v}
}

func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.topNode.val
	s.topNode = s.topNode.next
	return v
}

func (s *LinkedListStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.topNode.val
}

func (s *LinkedListStack) Flush() {
	s.topNode = nil
}

func (s *LinkedListStack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := s.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
