package main

import (
	"fmt"

	"awesomeProject/awesome"
)




//["LRUCache","put","put","get","get","put","get","get","get"]
//[[2],[2,1],[3,2],[3],[2],[4,3],[2],[3],[4]]
//输出：
//[null,null,null,2,1,null,-1,2,3]
//预期：
//[null,null,null,2,1,null,1,-1,3]

func main() {
	obj := awesome.Constructor(2);

	obj.Put(2,1);
	obj.Put(3,2);


	fmt.Sprintf("%+v", obj)
	val := obj.Get(3);
	val = obj.Get(2);

	obj.Put(4,3);

	fmt.Sprintf("%+v", obj)

	val = obj.Get(2);
	val = obj.Get(3);
	val = obj.Get(4);

	fmt.Sprintf("%+v", val)
	fmt.Sprintf("%+v", obj)
}
