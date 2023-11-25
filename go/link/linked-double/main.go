package main

import (
	"fmt"
)

type link[T any] struct {
	root   *elem[T]
	length int
}

func first[T any](l link[T]) *elem[T] {
	if l.length == 0 {
		return nil
	}
	return l.root
}

func last[T any](l link[T]) *elem[T] {
	if l.length == 0 {
		return nil
	}
	return l.root.prev
}

func add[T any](l link[T], value T) link[T] {
	if l.length == 0 {
		elem := &elem[T]{
			value: value,
		}
		elem.next = elem
		elem.prev = elem
		l.root = elem
	} else {
		elem := &elem[T]{
			value: value,
			prev:  last(l),
			next:  first(l),
		}
		last(l).next = elem
		first(l).prev = elem
	}
	l.length++
	return l
}

type elem[T any] struct {
	value T
	prev  *elem[T]
	next  *elem[T]
}

func main() {
	l := link[int]{}
	l = add(l, 1)
	l = add(l, 3)
	l = add(l, 5)
	l = add(l, 7)
	fmt.Println(first(l).value) // 1
	fmt.Println(last(l).value)  // 7

	f := first(l)
	fmt.Println(f.next.value)                     // 3
	fmt.Println(f.next.next.value)                // 5
	fmt.Println(f.next.next.next.next.value)      // 1
	fmt.Println(f.next.next.next.next.prev.value) // 7
}
