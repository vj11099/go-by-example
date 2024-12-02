package main

import "fmt"

func SlicesIndex[S ~[]E, E comparable](s S, e E) int {
	for i := range s {
		if e == s[i] {
			return i
		}
	}
	return -1
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head, tail *element[T]
}

func (lst *List[T]) Push(i T) {
	if lst.tail == nil {
		lst.tail = &element[T]{val: i}
		lst.head = lst.tail
	} else {
		lst.tail.next = &element[T]{val: i}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
	var items []T
	for i := lst.head; i != nil; i = i.next {
		items = append(items, i.val)
	}
	return items
}

func main() {
	s := []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	//explicitly calling generic functions
	//StringIndex of type [param1, param2]
	//where param1 of []string
	//and param2 of string
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
