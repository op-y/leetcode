package leetcode

import (
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	input := []int{4,3,2,7,8,2,3,1}
	//input := []int{1, 1, 2, 4}
	t.Logf("output: %v", findDisappearedNumbers(input))
}
