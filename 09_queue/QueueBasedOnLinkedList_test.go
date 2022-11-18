package _9_queue

import "testing"

func TestListQueue_EnQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	// head<-<-1<-2<-3<-4<-5<-6<-tail
	t.Log(q)
}

func TestListQueue_DeQueue(t *testing.T) {
	q := NewLinkedListQueue()
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	q.EnQueue(6)
	// head<-<-1<-2<-3<-4<-5<-6<-tail
	t.Log(q)
	q.DeQueue()
	// head<-<-2<-3<-4<-5<-6<-tail
	t.Log(q)
	q.DeQueue()
	// head<-<-3<-4<-5<-6<-tail
	t.Log(q)
	q.DeQueue()
	// head<-<-4<-5<-6<-tail
	t.Log(q)
	q.DeQueue()
	// head<-<-5<-6<-tail
	t.Log(q)
	q.DeQueue()
	// head<-<-6<-tail
	t.Log(q)
}
