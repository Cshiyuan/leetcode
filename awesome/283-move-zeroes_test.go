package awesome

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	//root := NewCodec().Deserialize("[1,2]")
	got := []int{0, 0, 1}
	moveZeroes(got)
	fmt.Println(got)
}
