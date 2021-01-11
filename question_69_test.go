package leetcode

import (
    "testing"
)

func TestMySqrt(t *testing.T) {
    a := 8
    t.Logf("%v", mySqrt(a))

    a = 4
    t.Logf("%v", mySqrt(a))

    a = 9
    t.Logf("%v", mySqrt(a))

    a = 11
    t.Logf("%v", mySqrt(a))

    a = 17
    t.Logf("%v", mySqrt(a))

    a = 2147395599
    t.Logf("%v", mySqrt(a))
}
