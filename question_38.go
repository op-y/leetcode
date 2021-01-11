package leetcode

import (
    "fmt"
)

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    num := "1"
    for i:=2; i<=n; i++ {
        s := ""
        d := ' '
        c := 0
        for _, digit := range num {
            if d == ' ' {
                d = digit
                c = 1
                continue
            }
            if digit == d {
                c++
            } else {
                s = fmt.Sprintf("%s%d%c", s, c, d)
                d = digit
                c = 1
            }
        }
        if c > 0 {
            s = fmt.Sprintf("%s%d%c", s, c, d)
        }
        num = s
    }
    return num
}
