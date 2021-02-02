package leetcode

import (
    "testing"
)

func TestThreeSum(t *testing.T) {
    input := []int{-1,0,1,2,-1,-4}
    t.Logf("input: %v", input)
    t.Logf("result: %v", threeSum(input))
}
