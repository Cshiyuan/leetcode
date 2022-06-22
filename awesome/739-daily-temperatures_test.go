package awesome

import (
	"fmt"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	got := dailyTemperatures([]int{73,74,75,71,69,72,76,73});
	fmt.Printf("got: %v", got)
}
