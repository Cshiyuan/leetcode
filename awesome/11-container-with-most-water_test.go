package awesome

import (
	"fmt"
	"testing"
)

func Test_maxArea(t *testing.T) {

	got := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(got)
}
