package mylist

import (
	"errors"
	"fmt"
)

const initialLen = 2

type List[T any] struct {
	arr      []T
	index    int
	length   int
	capasity int
}

func New[T any]() *List[T] {
	list := new(List[T])
	list.arr = make([]T, initialLen)
	list.index = 0
	list.length = 0
	list.capasity = initialLen
	return list
}

func (list *List[T]) Add(value T) {
	if list.length == list.capasity-1 {
		extend(list)
	}

	if list.length < list.capasity {
		list.arr[list.index] = value
		list.index++
		list.length++
	}
}

func (list *List[T]) Clear() {
	list.arr = make([]T, initialLen)
	list.capasity = initialLen
	list.index = 0
	list.length = 0
}

func (list *List[T]) RemoveAtIndex(index int) error {
	if index >= list.index {
		return errors.New("Index was out of bound")
	}
	for i := index; i < list.length; i++ {
		list.arr[i] = list.arr[i+1]
	}
	list.index--
	list.length--
	return nil

}
func (list *List[T]) RemoveAtFirst() error {
	if list.index == 0 {
		return errors.New("No element in list")
	}
	for i := 0; i < list.length; i++ {
		list.arr[i] = list.arr[i+1]
	}
	list.index--
	list.length--
	return nil
}
func (list *List[T]) RemoveAtLast() error {
	if list.index == 0 {
		return errors.New("No element in list")
	}
	list.arr[list.index-1] = list.arr[list.index]
	list.index--
	list.length--
	return nil
}

func extend[T any](list *List[T]) {
	newArr := make([]T, list.capasity*2)
	for i := 0; i < list.capasity; i++ {
		newArr[i] = list.arr[i]
	}
	list.arr = newArr
	list.capasity *= 2
}

func (list *List[T]) Len() int {
	return list.length
}

func (list *List[T]) Cap() int {
	return list.capasity
}

func (list *List[T]) String() string {
	var res string
	for i := 0; i < list.capasity; i++ {
		res += fmt.Sprintf("%v\n", list.arr[i])
	}
	return res
}
