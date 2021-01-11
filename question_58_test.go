package leetcode

import (
    "testing"
)

func TestLengthOfLastWord(t *testing.T) {
    input := "Hello World"
    t.Logf("%v", lengthOfLastWord(input))

    input = "Starbucks Delivers"
    t.Logf("%v", lengthOfLastWord(input))
}
