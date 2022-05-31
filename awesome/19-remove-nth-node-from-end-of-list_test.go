package awesome

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {

	got := removeNthFromEnd(buildListNode([]int{1, 2}), 2)
	fmt.Println(got)
}
