package awesome

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {

	got := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(got)
}
