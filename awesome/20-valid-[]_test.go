package awesome

import (
	"fmt"
	"testing"
)

func Test_isValid(t *testing.T) {

	got := isValid("(]")
	fmt.Println(got)
}
