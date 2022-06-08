package awesome

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	m := [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}
	rotate(m)
	fmt.Println(m)
}
