package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := New(6,1,3,5,2,4,7)
	// h.Push(6)
	// h.Push(1)
	// h.Push(3)
	// h.Push(5)
	// h.Push(2)
	// h.Push(4)
	// h.Push(7)
	fmt.Println("pop:", h.Pop())
	fmt.Println("pop:", h.Pop())
	fmt.Println("pop:", h.Pop())
	fmt.Println("pop:", h.Pop())
	fmt.Println("pop:", h.Pop())
	fmt.Println("pop:", h.Pop())
	fmt.Println("pop:", h.Pop())
	fmt.Println("pop:", h.Pop())
}