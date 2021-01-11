package leetcode

import (
    "testing"
)

func TestIsValid(t *testing.T) {
    input := "()"
    t.Logf("%t", isValid(input))

    input = "()[]{}"
    t.Logf("%t", isValid(input))

    input = "(]"
    t.Logf("%t", isValid(input))

    input = "([)]"
    t.Logf("%t", isValid(input))

    input = "{[]}"
    t.Logf("%t", isValid(input))

    input = "([]{}"
    t.Logf("%t", isValid(input))
}
