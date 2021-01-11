package leetcode

import (
    "testing"
)

func TestRob(t *testing.T) {
    input := []int{1,2,3,1}
    t.Logf("%v", rob(input))

    input = []int{2,7,9,3,1}
    t.Logf("%v", rob(input))

    input = []int{2,1,1,2}
    t.Logf("%v", rob(input))

    input = []int{1,2,1,1}
    t.Logf("%v", rob(input))

    input = []int{1,7,9,4}
    t.Logf("%v", rob(input))
}
