package leetcode

import (
    "testing"
)

func TestRomanToInt(t *testing.T) {
    input := "III"
    t.Logf("%v", romanToInt(input))

    input = "IV"
    t.Logf("%v", romanToInt(input))

    input = "IX"
    t.Logf("%v", romanToInt(input))

    input = "LVIII"
    t.Logf("%v", romanToInt(input))

    input = "MCMXCIV"
    t.Logf("%v", romanToInt(input))
}
