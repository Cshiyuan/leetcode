package awesome

import (
	"fmt"
	"testing"
)

func TestCodec_deserialize(t *testing.T) {


	got := NewCodec().Deserialize("[1,2,3,null,null,4,5]")
	fmt.Println(got)
}
