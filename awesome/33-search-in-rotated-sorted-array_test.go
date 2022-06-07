package awesome

import (
	"fmt"
	"testing"
)

func Test_search2(t *testing.T) {

	got := search([]int{4,5,6,7,0,1,2}, 0)
	fmt.Println(got)

}
