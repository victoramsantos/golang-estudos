package linkedlist

import "fmt"

type Something interface {
	GetValue() string
}

type Value struct {
	value string
}

func (value *Value) GetValue() string {
	return value.value
}

func NewValue(value string) *Value {
	return &Value{
		value: value,
	}
}

type StringLinkedList struct {
	next  *StringLinkedList
	value Something
}

func NewStringLinkedList(value Something) *StringLinkedList {
	return &StringLinkedList{
		next:  nil,
		value: value,
	}
}

func (list *StringLinkedList) Add(value Something) {
	if list.next != nil {
		list.next.Add(value)
		return
	}
	list.next = NewStringLinkedList(value)
}

func (list *StringLinkedList) ToString() string {
	var str string
	pointer := list
	for pointer != nil {
		str += fmt.Sprintf("%s,", pointer.value.GetValue())
		pointer = pointer.next
	}
	return fmt.Sprintf("[%s]", str)
}
