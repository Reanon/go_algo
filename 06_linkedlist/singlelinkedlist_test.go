package _6_linkedlist

import "testing"

func TestInsertToHead(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	// 10->9->8->7->6->5->4->3->2->1
	l.Print()
}

func TestInsertToTail(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	// 1->2->3->4->5->6->7->8->9->10
	l.Print()
}

func TestFindByIndex(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	// 返回的是结点的指针
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestDeleteNode(t *testing.T) {
	l := NewLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	// 1->2->3
	l.Print()
	//    singlelinkedlist_test.go:42: true
	t.Log(l.DeleteNode(l.head.next))
	// 2->3
	l.Print()
	//    singlelinkedlist_test.go:45: true
	t.Log(l.DeleteNode(l.head.next.next))
	// 2
	l.Print()
}
