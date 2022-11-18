package _1_array

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

/**
 * 1) 数组的插入、删除、按照下标随机访问操作；
 * 2）数组中的数据是int类型的；
 */

type Array struct {
	data []int      // 固定大小的数组，用满容量和满大小的切片来代替
	len  uint       // 真正长度
	cap  uint       // 容量
	lock sync.Mutex // 为了并发安全使用的锁
}

// NewArray 为数组初始化内存
func NewArray(length, capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	if length > capacity {
		log.Panic("len must be less than cap")
	}
	return &Array{
		data: make([]int, capacity, capacity),
		len:  0,
		cap:  capacity,
	}
}

// Len 返回真实长度
func (a *Array) Len() uint {
	return a.len
}

// Cap 返回容量
func (a *Array) Cap() uint {
	return a.len
}

// Append 增加一个元素
func (a *Array) Append(element int) {
	// 并发锁
	a.lock.Lock()
	defer a.lock.Unlock()
	// 大小等于容量，表示没多余位置了
	if a.len == a.cap {
		// 没容量，数组要扩容，扩容到两倍
		newCap := 2 * a.len
		// 如果之前的容量为0，那么新容量为1
		if a.cap == 0 {
			newCap = 1
		}
		newData := make([]int, newCap, newCap)
		// 把老数组的数据移动到新数组
		for k, v := range a.data {
			newData[k] = v
		}
		// 替换数组
		a.data = newData
		a.cap = newCap
	}
	// 把元素放在数组里
	a.data[a.len] = element
	// 真实长度+1
	a.len += 1
}

// AppendMany 增加多个元素
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

// isIndexOutOfRange 判断索引是否越界
func (a *Array) isIndexOutOfRange(index uint) bool {
	if index >= a.len {
		return true
	}
	return false
}

// Get 通过索引查找数组，索引范围[0,n-1]
func (a *Array) Get(index uint) (int, error) {
	if a.len == 0 || a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	return a.data[index], nil
}

// Insert 插入数值到索引index上
func (a *Array) Insert(index uint, v int) error {
	if a.Len() == uint(cap(a.data)) {
		return errors.New("full array")
	}
	if index != a.len && a.isIndexOutOfRange(index) {
		return errors.New("out of index range")
	}

	for i := a.len; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.len++
	return nil
}

func (a *Array) InsertToTail(v int) error {
	return a.Insert(a.Len(), v)
}

// Delete 删除索引index上的值
func (a *Array) Delete(index uint) (int, error) {
	if a.isIndexOutOfRange(index) {
		return 0, errors.New("out of index range")
	}
	v := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.len--
	return v, nil
}

// Print 打印数列
func (a *Array) Print() {
	var format string
	for i := uint(0); i < a.Len(); i++ {
		format += fmt.Sprintf("|%+v", a.data[i])
	}
	fmt.Println(format)
}
