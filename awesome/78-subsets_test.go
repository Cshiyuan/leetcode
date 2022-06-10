package awesome

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {

	got := subsets([]int{1,2,3});
	fmt.Print(got)
}
