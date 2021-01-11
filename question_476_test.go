package leetcode

import (
	"testing"
)

func TestFindComplement(t *testing.T) {
	input := 999999
	t.Logf("output: %v", findComplement(input))

	input = 5
	t.Logf("output: %v", findComplement(input))

	input = 1
	t.Logf("output: %v", findComplement(input))
}
