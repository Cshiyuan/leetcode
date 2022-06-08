package awesome

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {

	got := permute([]int{1,2,3});
	fmt.Println(got)
}
