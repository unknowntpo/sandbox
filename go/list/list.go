package main

type List struct {
	root *Element
	len  int
}

type Element struct {
	Value interface{}
	next  *Element
}

// New creates a new list
func New() *List {
	return &List{root: nil, len: 0}
}

// Len returns length of list
func (l *List) Len() int { return l.len }

func (l *List) Front() *Element { return l.root }

func (l *List) PushFront(value interface{}) {
	indirect := &l.root
	newNode := &Element{Value: value, next: l.root}
	*indirect = newNode
}
func (l *List) PushBack(value interface{}) {
	indirect := &l.root
	for ; *indirect != nil; indirect = &(*indirect).next {
		// do nothing
	}
	*indirect = &Element{Value: value, next: nil}
}

// Back returns the last element in the list
func (l *List) Back() *Element {
	e := l.Front()
	for ; e != nil && e.Next() != nil; e = e.Next() {
		// do nothing
	}
	return e
}

func (e *Element) Next() *Element {
	return e.next
}
