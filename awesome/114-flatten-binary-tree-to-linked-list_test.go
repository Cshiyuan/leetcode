package awesome

import (
	"fmt"
	"testing"
)

func Test_flatten(t *testing.T) {
	root := NewCodec().Deserialize("[1,2,5,3,4,null,6]")
	flatten(root)
	fmt.Print(root)
}
