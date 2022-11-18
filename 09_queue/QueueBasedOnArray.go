package _9_queue

import "fmt"

type ArrayQueue struct {
	// 使用空接口切片来保存数
	items    []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{make([]interface{}, n), n, 0, 0}
}

// EnQueue 入队
func (q *ArrayQueue) EnQueue(v interface{}) bool {
	// 当 tail == capacity 表示队列已满
	if q.tail == q.capacity {
		return false
	}
	q.items[q.tail] = v
	q.tail++
	return true
}

// DeQueue 出队
func (q *ArrayQueue) DeQueue() interface{} {
	if q.head == q.tail {
		return nil
	}
	v := q.items[q.head]
	q.head++
	return v
}

// String 通过字符串的方式输出
func (q *ArrayQueue) String() string {
	// 为空时
	if q.head == q.tail {
		return "empty items"
	}
	result := "head"
	for i := q.head; i <= q.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", q.items[i])
	}
	result += "<-tail"
	return result
}
