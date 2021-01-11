package leetcode

import (
    "testing"
)

func TestRemoveDuplications(t *testing.T) {
    input := []int{1, 1, 2}
    t.Logf("%v", input[0:removeDuplicates(input)])

    input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
    t.Logf("%v", input[0:removeDuplicates(input)])
}
