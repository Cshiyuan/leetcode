package awesome

import (
	"fmt"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {

	got := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(got)
}
