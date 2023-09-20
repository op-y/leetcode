package leetcode

import (
    "fmt"
)

func reOrderArray( array []int ) []int {
    if array == nil {
        return nil
    }
    n := len(array)
    if n <= 1 {
        return array
    }

    var odd int
    lastOdd := -1
    for i := 0; i < n; i++ {
        if array[i] % 2 == 1 {//找到一个奇数
            odd = i
            fmt.Printf("lastOdd %d, currentOdd %d\n", lastOdd, odd)
            if odd - lastOdd > 1 {//两个奇数不相邻
                for j := i; j > lastOdd + 1; j-- {
                    array[j], array[j - 1] = array[j - 1], array[j]
                }
                lastOdd++
            } else {
                lastOdd++
            }
        }
    }
    return array
}
