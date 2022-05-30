package awesome

import (
	"fmt"
	"math/bits"
)

func hammingDistance(x int, y int) int {

	return bits.OnesCount(uint(x ^ y))
}

func countBit(val int) int {

	str := fmt.Sprintf("%b", val)
	var count int

	for _, v := range str {
		if string(v) == "1" {
			count++
		}
	}
	return count
}
