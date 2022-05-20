package awesome

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	//[5,4,-1,7,8]  23
	//[-1,-2]

	//[-2,1,-3,4,-1,2,1,-5,4]
	got := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(got)
	//got, _, _ := lMaxSubArray([]int{-1,-1,-2,-2})
	//fmt.Println(got)
}
