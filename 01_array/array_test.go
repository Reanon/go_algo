package _1_array

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	capacity := 6
	arr := NewArray(0, uint(capacity))
	for i := 0; i < capacity-2; i++ {
		arr.Append(i)
		fmt.Print("cap ", arr.Cap(), " len ", arr.Len(), " array: ")
		arr.Print()
	}

	err := arr.Insert(uint(2), 999)
	if err != nil {
		t.Fatal(err.Error())
	}
	arr.Print()
	err = arr.InsertToTail(666)
	if err != nil {
		t.Fatal(err.Error())
	}
	arr.Print()
}

func TestDelete(t *testing.T) {
	capacity := 10
	arr := NewArray(0, uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	for i := 9; i >= 0; i-- {
		_, err := arr.Delete(uint(i))
		if nil != err {
			t.Fatal(err)
		}
		arr.Print()
	}
}

func TestGet(t *testing.T) {
	capacity := 10
	arr := NewArray(0, uint(capacity))
	for i := 0; i < capacity; i++ {
		err := arr.Insert(uint(i), i+1)
		if nil != err {
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.Get(0))
	t.Log(arr.Get(9))
	t.Log(arr.Get(11))
}
