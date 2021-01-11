package leetcode

import (
    "testing"
)

func TestLongestCommonPrefix(t *testing.T) {
    input := []string{"flower","flow","flight"}
    t.Logf("%s", longestCommonPrefix(input))

    input = []string{"dog", "racecar", "car"}
    t.Logf("%s", longestCommonPrefix(input))

    input = []string{"world", "world"}
    t.Logf("%s", longestCommonPrefix(input))

    input = []string{"", ""}
    t.Logf("%s", longestCommonPrefix(input))

    input = []string{}
    t.Logf("%s", longestCommonPrefix(input))

    input = []string{"hello"}
    t.Logf("%s", longestCommonPrefix(input))
}
