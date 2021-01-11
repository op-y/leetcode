package leetcode

import (
    "testing"
)

func TestPlusOne(t *testing.T) {
    input := []int{1,2,3}
    t.Logf("%v", plusOne(input))

    input = []int{4,3,2,1}
    t.Logf("%v", plusOne(input))

    input = []int{9,9,9,9}
    t.Logf("%v", plusOne(input))
}
