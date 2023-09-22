package leetcode

import (
	"testing"
)

func TestTopKstrings692(t *testing.T) {
	input := []string{"1", "2", "3", "4"}
	k := 2
	t.Logf("result: %v", topKstrings(input, k))
}
