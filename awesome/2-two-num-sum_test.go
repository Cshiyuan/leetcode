package awesome

import (
	"fmt"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {


//[9,9,9,9,9,9,9]
//	[9,9,9,9]
	got := addTwoNumbers(buildListNode([]int{9,9,9,9,9,9,9}), buildListNode([]int{9,9,9,9}))
	fmt.Println(got)
}
