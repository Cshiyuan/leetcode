package awesome

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome2(head *ListNode) bool {

	if head.Next == nil {
		return true
	}


	totalNum := 0

	flag := true

	// 遍历转成数组
	for ;; {
		if flag {
			totalNum = totalNum + head.Val
		} else {
			totalNum = totalNum - head.Val
		}
		if head.Next == nil {
			break
		}
		head = head.Next
		flag = !flag
	}
	return totalNum == 0
}

func isPalindrome(head *ListNode) bool {

	// 只有单个节点
	if head.Next == nil {
		return true
	}

	vals := make([]int, 0 ,0)

	// 遍历转成数组
	for ;; {
		vals = append(vals, head.Val)
		if head.Next == nil {
			break
		}
		head = head.Next
	}

	length := len(vals)

	secondPos := 0
	firstPos := 0

	// 不是偶数
	if length % 2 != 0 {
		t := length - 1
		secondPos = (t / 2) + 1
		firstPos = (t / 2) - 1
	} else {
		secondPos = length / 2
		firstPos = secondPos - 1
	}

	b := true

	for ;; {
		if secondPos > length - 1  || firstPos < 0{
			break
		}
		if vals[firstPos] != vals[secondPos] {
			b = false
			break
		}
		secondPos ++
		firstPos --
	}

	return b
}
