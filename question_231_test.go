package leetcode

import (
    "testing"
)

func TestIsPowerOfTwo(t *testing.T) {
    a := 16
    t.Logf("%v", isPowerOfTwo(a))

    a = 218
    t.Logf("%v", isPowerOfTwo(a))

    a = 1
    t.Logf("%v", isPowerOfTwo(a))

    a = 0
    t.Logf("%v", isPowerOfTwo(a))
}
