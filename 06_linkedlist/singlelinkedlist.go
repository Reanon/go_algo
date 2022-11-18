package _6_linkedlist

import "fmt"

/*
  单链表基本操作
*/

// ListNode 单链表结点
type ListNode struct {
	next  *ListNode
	value interface{}
}

// LinkedList 单向链表
type LinkedList struct {
	head   *ListNode
	length uint
}

// NewListNode 新建一个结点
func NewListNode(v interface{}) *ListNode {
	// 使用取地址的方式初始化结构体
	return &ListNode{nil, v}
}

// GetNext 获取下一个结点
func (n *ListNode) GetNext() *ListNode {
	return n.next
}

// GetValue 获取结点的值
func (n *ListNode) GetValue() interface{} {
	return n.value
}

// NewLinkedList 新建一个单向链表
func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

// InsertAfter 在某个节点后面插入节点
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	l.length++
	return true
}

// InsertBefore 在某个节点前面插入节点
func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == l.head {
		return false
	}
	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	l.length++
	return true
}

// InsertToHead 在链表头部插入节点
func (l *LinkedList) InsertToHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// InsertToTail 在链表尾部插入节点
func (l *LinkedList) InsertToTail(v interface{}) bool {
	cur := l.head
	for nil != cur.next {
		cur = cur.next
	}
	return l.InsertAfter(cur, v)
}

// FindByIndex 通过索引查找节点
func (l *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= l.length {
		return nil
	}
	cur := l.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// DeleteNode 删除传入的节点
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := l.head.next
	pre := l.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	l.length--
	return true
}

// Print 打印链表
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
