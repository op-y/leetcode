package leetcode

import (
    "testing"
)

func TestAddBinary(t *testing.T) {
    a := "11"
    b := "1"
    t.Logf("%v", addBinary(a, b))

    a = "1010"
    b = "1011"
    t.Logf("%v", addBinary(a, b))
}
