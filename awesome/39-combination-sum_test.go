package awesome

import (
	"fmt"
	"testing"
)

func Test_combinationSum(t *testing.T) {

	got := combinationSum([]int{2,3,5}, 8)
	fmt.Println(got)
}
