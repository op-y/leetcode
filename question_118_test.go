package leetcode

import (
    "testing"
)

func TestGenerate(t *testing.T) {
    result := generate(0)
    t.Logf("%v", result)

    result = generate(1)
    t.Logf("%v", result)

    result = generate(2)
    t.Logf("%v", result)

    result = generate(3)
    t.Logf("%v", result)
}
