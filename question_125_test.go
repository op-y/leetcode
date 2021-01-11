package leetcode

import (
    "testing"
)

func TestIsPalindrome(t *testing.T) {
    input := "A man, a plan, a canal: Panama"
    t.Logf("%v", isPalindrome(input))

    input = "race a car"
    t.Logf("%v", isPalindrome(input))

    input = "0P"
    t.Logf("%v", isPalindrome(input))
}
