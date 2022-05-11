package awesome

import (
	"fmt"
	"testing"
)

func Test_findKthLargest(t *testing.T) {

		num := findKthLargest([]int{3,2,1,5,6,4}, 2)
		fmt.Sprintf("num :%v", num)
}
