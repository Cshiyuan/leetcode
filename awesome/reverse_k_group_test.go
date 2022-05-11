package awesome

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {

	reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}, 2)

}
