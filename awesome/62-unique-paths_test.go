package awesome

import (
	"fmt"
	"testing"
)

func Test_uniquePaths(t *testing.T) {

	got := uniquePaths(3, 7);
	fmt.Println(got)
}
