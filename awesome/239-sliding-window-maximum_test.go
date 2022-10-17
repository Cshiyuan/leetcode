package awesome

import (
	"fmt"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {

	got := maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3);
	fmt.Printf("got: %v", got)
}
