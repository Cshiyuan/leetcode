package main

import (
	"fmt"
	"testing"
)

func Test_fastQuick(t *testing.T) {

	got := fastQuick([]int{4, 1, 6, 2, 3, 5})
	fmt.Printf("got: %v", got)
}
