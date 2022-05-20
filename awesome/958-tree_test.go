package awesome

import (
	"fmt"
	"testing"
)
//[2,null,4,9,8,null,null,4]
func Test_isCompleteTree(t *testing.T) {

	root := NewCodec().Deserialize("[1,2,3,null,null,7,8]")
	got := isCompleteTree(root)
	fmt.Println(got)


	//[]*int{2,nil,4,9,8,nil,nil,4}
}
