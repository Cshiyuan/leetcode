package awesome

import (
	"fmt"
	"testing"
)

//[3,9,20,null,null,15,7]
func Test_levelOrder(t *testing.T) {

	got := levelOrder(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
			},
			Right: &TreeNode{
				Val:   7,
			},
		},

	})
	fmt.Println(got)
}
