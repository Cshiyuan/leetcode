package awesome

import (
	"fmt"
	"testing"
)

func Test_searchRange(t *testing.T) {
	got := searchRange([]int{5,7,7,8,8,10}, 8)
	fmt.Println(got)
}
