package leetcode

import (
    "testing"
)

func TestIsIsomorphic(t *testing.T) {
    a := "add"
    b := "egg"
    t.Logf("%v", isIsomorphic(a, b))

    a = "foo"
    b = "bar"
    t.Logf("%v", isIsomorphic(a, b))

    a = "ab"
    b = "aa"
    t.Logf("%v", isIsomorphic(a, b))
}
