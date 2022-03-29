package linkedlist

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	next  *LinkedList
	value int
}

func NewLinkedList(value int) *LinkedList {
	return &LinkedList{
		next:  nil,
		value: value,
	}
}

func (list *LinkedList) Add(value int) {
	if list.next != nil {
		list.next.Add(value)
		return
	}
	list.next = NewLinkedList(value)
}

func (list *LinkedList) GetValueAt(position int) (int, error) {
	if position == 0 {
		return list.value, nil
	}
	if position < 0 {
		return 0, errors.New("InvalidPosition")
	}
	if list.next == nil {
		return 0, errors.New("IndexOutOfBound")
	}
	return list.next.GetValueAt(position - 1)
}

func (list *LinkedList) ToString() string {
	var str string
	pointer := list
	for pointer != nil {
		str += fmt.Sprintf("%d,", pointer.value)
		pointer = pointer.next
	}
	return fmt.Sprintf("[%s]", str)
}

func (list *LinkedList) Length() int {
	if list.next != nil {
		return list.next.Length() + 1
	}
	return 1
}

func (list *LinkedList) AddAt(value int, position int) error {
	if position == 0 {
		temp := NewLinkedList(list.value)
		temp.next = list.next
		list.value = value
		list.next = temp
		return nil
	}
	if position < 0 {
		return errors.New("InvalidPosition")
	}
	if list.next == nil {
		return errors.New("IndexOutOfBound")
	}
	return list.next.AddAt(value, position-1)
}

// func (list *LinkedList) RemoveAt(position int) (int, error) {
// }
