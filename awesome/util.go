package awesome


// 辅助工具，进行链表的创建
func buildListNode(nums []int) *ListNode {
	root := &ListNode{}
	curNode := root

	for i, n := range nums {
		curNode.Val = n

		if !(i == len(nums) - 1) {
			curNode.Next = &ListNode{}
		}
		curNode = curNode.Next
	}
	return root
}