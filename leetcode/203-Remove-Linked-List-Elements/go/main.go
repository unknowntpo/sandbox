package main

import "fmt"

type List struct {
	Head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func New() *List {

}

func (l *List) PushBack(val int) *ListNode {

}

func (l *List) RemoveElements(val int) {
	l.head = removeElements(l.head, val)
}

func removeElements(head *ListNode, val int) *ListNode {

}
