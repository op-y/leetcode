package leetcode

import (
    "testing"
)

func TestIsPalindrome(t *testing.T) {
    input := 121
    t.Logf("%v", isPalindrome(input))

    input = -121
    t.Logf("%v", isPalindrome(input))

    input = 10
    t.Logf("%v", isPalindrome(input))

    input = 0
    t.Logf("%v", isPalindrome(input))
}
