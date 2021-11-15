package main

import "testing"

func TestremoveElements(t *testing.T) {
	t.Run("remove multiple nodes and remove last elements", func(t *testing.T) {
		l := New()
		elems := []int{1, 2, 6, 3, 4, 5, 6}
		target := 6
		for _, e := range elems {
			l.PushBack(e)
		}
		l.Head = removeElements(l.Head, target)
		assert.Equal(t, "[1,2,3,4,5]", l.String())
	})
}
