package leetcode

import (
    "testing"
)

func TestThirdMax(t *testing.T) {
    input := []int{1,2}
    t.Logf("isPalindrome: %v", thirdMax(input))

    input = []int{3,2,1}
    t.Logf("isPalindrome: %v", thirdMax(input))

    input = []int{2,2,3,1}
    t.Logf("isPalindrome: %v", thirdMax(input))

    input = []int{1,1,2}
    t.Logf("isPalindrome: %v", thirdMax(input))

    input = []int{1,2,2}
    t.Logf("isPalindrome: %v", thirdMax(input))
}
