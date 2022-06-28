package awesome

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	list := []int{2,44,4,8,33,1,22,-11,6,34,55,54,9}
	QuickSort(list,0,len(list)-1)
	fmt.Printf("%v", list)
}
