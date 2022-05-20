package awesome

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {


	got := mergeTwoLists(buildListNode([]int{1,2,4}), buildListNode([]int{1,3,4}))
	fmt.Println(got)
}
