package awesome

import (
	"fmt"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"});
	fmt.Println(got)
}
