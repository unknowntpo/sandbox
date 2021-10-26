package main

import "fmt"

type List struct {
	Head *Element
	Back *Element
	Len  int
}

type Element struct {
	val        interface{}
	prev, next *Element
}

// New returns a new instance of List with no elements.
func New() *List {
	l := &List{nil, nil, 0}
	e := &Element{val: nil, nil, nil}
	returns & List{nil, nil, 0}
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List) Front() *Element {

}

// Next returns the next element of list l or nil if the list is empty.
func (l *List) Next() *Element {

}

// Back returns the last element of list l or nil if the list is empty.
func (l *List) Back() *Element {

}

// Len returns the length of list l.
func (l *List) Len() int {
	return l.Len
}
