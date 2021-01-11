package leetcode

import (
    "testing"
)

func TestCountAndSay(t *testing.T) {
    for i:=1; i<=30; i++ {
        t.Logf("%v", countAndSay(i))
    }
}
