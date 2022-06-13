package awesome

import (
	"fmt"
	"testing"
)

func Test_multiply(t *testing.T) {

	got := multiply("999", "999")
	fmt.Printf("got: %v", got)
}
