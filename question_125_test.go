package leetcode

import (
	"testing"
)

func TestIsPalindrome125(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	t.Logf("%v", isPalindrome125(input))

	input = "race a car"
	t.Logf("%v", isPalindrome125(input))

	input = "0P"
	t.Logf("%v", isPalindrome125(input))
}
