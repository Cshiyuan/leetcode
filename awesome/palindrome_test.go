package awesome

import "testing"

func Test_isPalindrome(t *testing.T) {
	isPalindrome(&ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
	})
}
