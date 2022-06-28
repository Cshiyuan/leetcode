package awesome

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {

	sort.Slice(people, func(i, j int)bool{
		a,b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	fmt.Printf("people: %v \n", people)
	ans := make([][]int, len(people))

	// 填空的思想
	for _, person := range people {
		spaces := person[1] + 1
		for i := range ans {

			if ans[i] == nil {
				spaces --
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}
	return ans
}