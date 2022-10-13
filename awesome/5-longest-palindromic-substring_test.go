package awesome

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {

	got := longestPalindrome("babad")

	got = longestPalindrome_V2("babad")
	fmt.Println(got)
}
