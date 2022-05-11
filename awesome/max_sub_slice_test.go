package awesome

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	//[5,4,-1,7,8]  23
//[-1,-2]

	got, _, _ := lMaxSubArray([]int{-1,-1,-2,-2})
	fmt.Println(got)
}
