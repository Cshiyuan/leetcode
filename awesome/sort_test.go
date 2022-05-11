package awesome

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	numbers := []int{5,4,20,3,8,2,8}
	quickSort(numbers,0,len(numbers)-1)
	fmt.Println(numbers)
}
