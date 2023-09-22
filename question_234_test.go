package leetcode

import (
	"testing"
)

func TestMergeTwoLists234(t *testing.T) {
	l2 := new(ListNode)
	l2.Val = 1

	l22 := new(ListNode)
	l22.Val = 3
	l2.Next = l22

	l23 := new(ListNode)
	l23.Val = 2
	l22.Next = l23

	l24 := new(ListNode)
	l24.Val = 4
	l23.Next = l24

	l25 := new(ListNode)
	l25.Val = 3
	l24.Next = l25

	l26 := new(ListNode)
	l26.Val = 2
	l25.Next = l26

	l27 := new(ListNode)
	l27.Val = 1
	l26.Next = l27

	//t.Logf("isPalindrome: %v", isPalindrome(l1))

	t.Logf("isPalindrome: %v", isPalindrome(l2))
}
