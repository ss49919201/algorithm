package main

import "fmt"

type link[T any] struct {
	elems []*elem[T]
}

func length[T any](l link[T]) int {
	return len(l.elems)
}

func first[T any](l link[T]) *elem[T] {
	if length(l) == 0 {
		return nil
	}
	return l.elems[0]
}

func last[T any](l link[T]) *elem[T] {
	if length(l) == 0 {
		return nil
	}
	return l.elems[length(l)-1]
}

func add[T any](l link[T], value T) link[T] {
	if length(l) == 0 {
		elem := &elem[T]{
			value: value,
		}
		elem.next = elem
		elem.prev = elem
		l.elems = append(l.elems, elem)
	} else {
		elem := &elem[T]{
			value: value,
			prev:  l.elems[len(l.elems)-1],
			next:  l.elems[0],
		}
		last(l).next = elem
		first(l).prev = elem

		l.elems = append(l.elems, elem)
	}
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
