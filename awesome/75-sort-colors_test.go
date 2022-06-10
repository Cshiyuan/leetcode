package awesome

import (
	"fmt"
	"testing"
)

func Test_sortColors(t *testing.T) {
	list := []int{2,0,2,1,1,0}
	sortColors(list)
	fmt.Print(list)
}
