package awesome

import (
	"fmt"
	"testing"
)
//[3,9,20,15,7]
//[9,3,15,20,7]
func Test_buildTree(t *testing.T) {

	got := buildTree([]int{1,2,3}, []int{3,2,1});
	fmt.Print(got)
}
