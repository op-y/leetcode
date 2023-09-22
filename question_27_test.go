package leetcode

import (
	"testing"
)

func TestRemoveDuplications27(t *testing.T) {
	input := []int{3, 2, 2, 3}
	t.Logf("%v", input[0:removeElement(input, 3)])

	input = []int{0, 1, 2, 2, 3, 0, 4, 2}
	t.Logf("%v", input[0:removeElement(input, 2)])
}
