package _9_queue

import "fmt"

type CircularQueue struct {
	items    []interface{}
	capacity int
	head     int
	tail     int
}

func NewCircularQueue(n int) *CircularQueue {
	if n == 0 {
		return nil
	}
	return &CircularQueue{make([]interface{}, n), n, 0, 0}
}

func (q *CircularQueue) IsEmpty() bool {
	// 栈空条件：head==tail为true
	if q.head == q.tail {
		return true
	}
	return false
}

func (q *CircularQueue) IsFull() bool {
	// 栈满条件：(tail + 1) % capacity == head 为true
	if q.head == (q.tail+1)%q.capacity {
		return true
	}
	return false
}

func (q *CircularQueue) EnQueue(v interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.items[q.tail] = v
	q.tail = (q.tail + 1) % q.capacity
	return true
}

func (q *CircularQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.items[q.head]
	q.head = (q.head + 1) % q.capacity
	return v
}

func (q *CircularQueue) String() string {
	if q.IsEmpty() {
		return "empty items"
	}
	result := "head"
	var i = q.head
	for true {
		result += fmt.Sprintf("<-%+v", q.items[i])
		i = (i + 1) % q.capacity
		if i == q.tail {
			break
		}
	}
	result += "<-tail"
	return result
}
