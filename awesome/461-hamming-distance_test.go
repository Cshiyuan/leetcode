package awesome

import (
	"fmt"
	"testing"
)

func Test_countBit(t *testing.T) {

	got := hammingDistance(1, 4)
	fmt.Println(got)

	got = countBit(3)
	fmt.Println(got)

}
