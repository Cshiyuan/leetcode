package awesome



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	count := 0
	l1CurNode := l1
	l2CurNode := l2

	// 需要返回的结果
	resultRoot := &ListNode{}
	curResultNode := resultRoot
	preVal := 0
	for ; ; {


		var (
			l1var = 0
			l2var = 0
		)
		if l1CurNode != nil {
			l1var = l1CurNode.Val
			l1CurNode = l1CurNode.Next
		}
		if l2CurNode != nil {
			l2var = l2CurNode.Val
			l2CurNode = l2CurNode.Next
		}

		// 当前位 + 进位
		resultVal := l1var + l2var + preVal
		if resultVal >= 10 {
			preVal = 1 //记录进位
			resultVal = resultVal - 10 //记录当前位置
		} else {
			preVal = 0
		}


		curResultNode.Val = resultVal
		if l1CurNode == nil && l2CurNode == nil && preVal == 0 {
			break
		}
		curResultNode.Next = &ListNode{}
		curResultNode = curResultNode.Next

		count++
	}
	return resultRoot
}
