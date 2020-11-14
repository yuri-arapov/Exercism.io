package listops

//package main

import "fmt"

// Single-linked list of ints.
type IntList struct {
	data int
	next *IntList
}

// Return true if lists s1 and s2 contain the same data element-wise.
func Equal(s1, s2 *IntList) bool {
	for s1 != nil && s2 != nil && s1.data == s2.data {
		s1 = s1.next
		s2 = s2.next
	}
	return s1 == nil && s2 == nil
}

// Turn list into string.
func (s *IntList) String() (res string) {
	res = "["
	for sep := ""; s != nil; s = s.next {
		res = res + sep + fmt.Sprintf("%d", s.data)
		sep = "->"
	}
	return res + "]"
}

// Make list from slice of ints.
func MakeList(data ...int) *IntList {
	if len(data) == 0 {
		return nil
	}
	head := &IntList{data: data[0], next: nil}
	tail := head
	for _, d := range data[1:] {
		tail.next = &IntList{data: d, next: nil}
		tail = tail.next
	}
	return head
}

// Return length of the list.
func (s *IntList) Length() (res int) {
	for ; s != nil; s = s.next {
		res++
	}
	return res
}

// Return last element of the list.
func (s *IntList) Tail() *IntList {
	if s == nil {
		return nil
	}
	for ; s.next != nil; s = s.next {
	}
	return s
}

// Append list s2 to the end of list s1, return combined list.
// Destructive for s1.
func (s1 *IntList) Append(s2 *IntList) *IntList {
	if s1 == nil {
		return s2
	}
	if s2 == nil {
		return s1
	}
	s1.Tail().next = s2
	return s1
}

// Reverse the list. Non-destructive.
func (s *IntList) Reverse() *IntList {
	var head *IntList = nil
	for ; s != nil; s = s.next {
		current := *s
		current.next = head
		head = &current
	}
	return head
}

// Return list of elements that given predicate f returns true for.
// Non-destructive.
func (s *IntList) Filter(f func(int) bool) *IntList {
	var head *IntList = nil
	for ; s != nil; s = s.next {
		if f(s.data) {
			e := *s
			e.next = head
			head = &e
		}
	}
	return head.Reverse()
}

// Return list each element of which is produced from initial list
// by applying function f. Non-destructive.
func (s *IntList) Map(f func(int) int) *IntList {
	var head *IntList = nil
	for ; s != nil; s = s.next {
		e := &IntList{data: f(s.data), next: head}
		head = e
	}
	return head.Reverse()
}

// Concatenate lists. Destructive.
func (s *IntList) Concat(lists ...*IntList) *IntList {
	res := s
	for _, l := range lists {
		res = res.Append(l)
	}
	return res
}

// Fold left to right.
func (s *IntList) Foldl(fn func(acc, e int) int, init int) int {
	acc := init
	for ; s != nil; s = s.next {
		acc = fn(acc, s.data)
	}
	return acc
}

// Fold right to left
func (s *IntList) Foldr(fn func(e, acc int) int, init int) int {
	return s.Reverse().Foldl(
		func(acc, e int) int { return fn(e, acc) }, // swap arguments
		init)
}
