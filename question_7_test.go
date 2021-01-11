package leetcode

import (
    "testing"
)

func TestReverse(t *testing.T) {
    input := int32(123)
    output := reverse(input)
    t.Logf("%v", output)

    input = int32(-123)
    output = reverse(input)
    t.Logf("%v", output)

    input = int32(120)
    output = reverse(input)
    t.Logf("%v", output)

    input = int32(0)
    output = reverse(input)
    t.Logf("%v", output)

    input = int32(2147483647)
    output = reverse(input)
    t.Logf("%v", output)

    input = int32(-2147483648)
    output = reverse(input)
    t.Logf("%v", output)
}
