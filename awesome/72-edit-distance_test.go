package awesome

import (
	"fmt"
	"testing"
)

func Test_minDistance(t *testing.T) {

	got := minDistance("horse", "ros")
	fmt.Printf("got: %v", got)
}
