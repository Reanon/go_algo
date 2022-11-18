package _8_stack

import "fmt"

/*
基于数组实现的栈
*/

type ArrayStack struct {
	// 数据
	data []interface{}
	// 栈顶, 大于等于0时有值
	top int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		// 设置栈的长度为0，容量为 32
		data: make([]interface{}, 0, 32),
		// 初始化栈顶为 -1
		top: -1,
	}
}

func (s *ArrayStack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *ArrayStack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top += 1
	}

	if s.top > len(s.data)-1 {
		// append 会自动扩容
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

func (s *ArrayStack) Pop() interface{} {
	// 弹出栈前，先判断是否为空
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top -= 1
	return v
}

func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

func (s *ArrayStack) Flush() {
	s.top = -1
}

func (s *ArrayStack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := s.top; i >= 0; i-- {
			fmt.Println(s.data[i])
		}
	}
}
