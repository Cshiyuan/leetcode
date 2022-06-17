package awesome

import (
	"fmt"
	"testing"
)

func Test_maxDepth(t *testing.T) {

	root := NewCodec().Deserialize("[1,null,2,3]")
	got := maxDepth(root)
	fmt.Println(got)
}
