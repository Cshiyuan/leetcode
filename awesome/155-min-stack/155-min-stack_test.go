package _155_min_stack

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {

	got := Constructor()
	got.Push(-2)
	got.Push(0)
	got.Push(-3)
	val := got.GetMin()
	fmt.Println(val)
	got.Pop()
	got.Top()
	got.GetMin()
	fmt.Println(val)
}
