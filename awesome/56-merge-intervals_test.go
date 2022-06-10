package awesome

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {


	got := merge([][]int{{1,4},{2,3}})
	fmt.Println(got)
}
