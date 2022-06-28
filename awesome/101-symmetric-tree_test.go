package awesome

import (
	"fmt"
	"testing"
)

func Test_isSymmetric(t *testing.T) {

	root := NewCodec().Deserialize("[1,null,2,null,3,null,4,null,5]")
	got := isSymmetric(root)
	fmt.Println(got)
}
