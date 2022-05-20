package awesome

import (
	"math"
	"strconv"
)

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	l1num := buildNum(l1)
	l2num := buildNum(l2)

	return buildTree(l1num + l2num)
}

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

func buildNum(l *ListNode) int {

	var num int
	var count int
	for ; ; {
		if l == nil {
			break
		}

		num = num + l.Val*int(math.Pow10(count))
		l = l.Next
		count++
	}
	return num
}

func buildTree(num int) *ListNode {
	//var num int
	numStr := strconv.Itoa(num)

	root := &ListNode{
		Val:  0,
		Next: nil,
	}
	curNode := root
	for i := len(numStr) - 1; i >= 0; i-- {

		num, _ := strconv.Atoi(string(numStr[i]))
		curNode.Val = num
		if i != 0 {
			curNode.Next = &ListNode{}
			curNode = curNode.Next
		}

	}
	return root
}
