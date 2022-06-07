package awesome

import (
	"fmt"
	"testing"
)

func Test_combinationSum(t *testing.T) {

	got := combinationSum([]int{2,3,6,7}, 7)
	fmt.Println(got)
}
