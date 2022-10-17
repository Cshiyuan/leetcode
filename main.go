package main

import (
	"fmt"
)




//["LRUCache","put","put","get","get","put","get","get","get"]
//[[2],[2,1],[3,2],[3],[2],[4,3],[2],[3],[4]]
//输出：
//[null,null,null,2,1,null,-1,2,3]
//预期：
//[null,null,null,2,1,null,1,-1,3]

func update(s []int) {
	s[1] = 10
	fmt.Printf("%p\n", s)
}

func main() {
	a := []int{0, 1, 2, 3, 4}
	fmt.Println(a)
	fmt.Printf("%p\n", a)
	update(a)
	fmt.Println(a)
}
