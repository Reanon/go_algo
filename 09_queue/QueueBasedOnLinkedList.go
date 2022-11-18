package _9_queue

import "fmt"

// ListNode 单链表结点
type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	// 初始化链表队列
	return &LinkedListQueue{nil, nil, 0}
}

func (q *LinkedListQueue) EnQueue(v interface{}) {
	node := &ListNode{v, nil}
	if nil == q.tail {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.length++
}

func (q *LinkedListQueue) DeQueue() interface{} {
	if q.head == nil {
		return nil
	}
	v := q.head.val
	q.head = q.head.next
	q.length--
	return v
}

func (q *LinkedListQueue) String() string {
	if q.head == nil {
		return "empty items"
	}
	result := "head<-"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
