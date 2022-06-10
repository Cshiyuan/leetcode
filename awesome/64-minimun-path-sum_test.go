package awesome

import (
	"fmt"
	"testing"
)

func Test_minPathSum(t *testing.T) {

	got := minPathSum([][]int{{1,3,1},{1,5,1},{4,2,1}});
	fmt.Println(got)
}
