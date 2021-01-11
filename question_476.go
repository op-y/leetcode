package leetcode

import (
    "fmt"
    //"unsafe"
)

func findComplement(num int) int {
    fmt.Printf("%b\n", num)
    count := 0
    n := num
    for n != 0 {
        count++
        n = n >> 1
    }
    fmt.Printf("Bit Count: %d\n", count)

    mask := 0x7fffffffffffffff
    mask = mask << count
    fmt.Printf("Mask: %b\n", mask)
    mask = ^mask
    fmt.Printf("Mask: %b\n", mask)

    result := (^num) & mask
    fmt.Printf("Result: %b\n", result)
    return result
}
