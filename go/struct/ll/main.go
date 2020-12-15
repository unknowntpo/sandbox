// This program implement basic singly linked list operation
// at 2020q3 quiz1
// Ref:
// 	- [2020q3 Homework1 (quiz1)](https://hackmd.io/@unknowntpo/quiz1_ans)

package main

import "fmt"

// add_entry
// find_entry
// remove_entry
// swap pair

type Element struct {
	Value int
	Next  *Element
}
type List struct {
	head *Element
}

func newNode(value int) *Element {
	return &Element{Value: value, Next: nil}
}

// CreateList will create an empty list (head == nil)
func CreateList() *List {
	return &List{nil}
}
func (l *List) InsertHead(value int) {
	if l.head == nil {
		l.head = &Element{Value: value, Next: nil}
		return
	}
	l.head = l.head.insertHead(value)
}
func (e *Element) insertHead(value int) *Element {
	return &Element{Value: value, Next: e}
}
func (l *List) Clear()          { l.head = nil }
func (e *Element) findEntry()   {}
func (e *Element) removeEntry() {}
func (e *Element) swap_pair()   {}
func (l *List) PrintList() {
	if l == nil {
		fmt.Println("nil list")
		return
	}
	l.head.printList()
}
func (e *Element) printList() {
	// Declare a pointer to head of head of linked list
	h := e
	for h != nil {
		fmt.Print(h.Value, " -> ")
		h = h.Next
	}
	fmt.Println("nil")
}

func main() {
	in := []int{5, 4, 3, 2, 1}

	fmt.Println("Create new list...")
	list := CreateList()

	fmt.Println("Insert element into list...")
	for _, v := range in {
		list.InsertHead(v)
	}

	fmt.Println("Print the list...")
	list.PrintList()
	fmt.Println("Clear the list...")
	list.Clear()
	list.PrintList()
}
