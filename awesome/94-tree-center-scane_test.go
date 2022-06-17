package awesome

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {

	root := NewCodec().Deserialize("[1,2,2,3,4,4,3]")
	got := inorderTraversal(root)
	fmt.Println(got)
}
