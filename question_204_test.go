package leetcode

import (
    "testing"
)

func TestCountPrimes(t *testing.T) {
    input := 10
    t.Logf("%v", countPrimes(input))

    input = 100
    t.Logf("%v", countPrimes(input))

    input = 1000
    t.Logf("%v", countPrimes(input))

    input = 10000
    t.Logf("%v", countPrimes(input))

    input = 100000
    t.Logf("%v", countPrimes(input))

    input = 1000000
    t.Logf("%v", countPrimes(input))

    input = 10000000
    t.Logf("%v", countPrimes(input))

    input = 100000000
    t.Logf("%v", countPrimes(input))

    input = 100000000
    t.Logf("%v", countPrimes(input))
}
