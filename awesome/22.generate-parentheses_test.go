package awesome

import (
	"fmt"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	got := generateParenthesis(3);

	fmt.Println(got)
}
