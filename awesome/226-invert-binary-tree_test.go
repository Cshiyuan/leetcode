package awesome

import (
	"fmt"
	"testing"
)

func Test_invertTree(t *testing.T) {

	root := NewCodec().Deserialize("[1,2]")
	got := invertTree(root)
	fmt.Println(got)
}
