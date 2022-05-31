package awesome

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {

	got := longestPalindrome("babad")
	fmt.Println(got)
}
