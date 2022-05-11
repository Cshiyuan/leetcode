package awesome

import (
	"fmt"
	"testing"
)

func Test_threeSum(t *testing.T) {

	r := threeSum([]int{0,0,0,1,1,1,1,-1,-1,-1,-1,-1})

	fmt.Println(r)
}
