package awesome

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
//[186,419,83,408]
//	6249
//[3,7,405,436]
//	8839

	//输入：coins = [1, 2, 5], amount = 11
	got := coinChange2([]int{3,7,405,436},  8839)
	fmt.Println(got)
}
