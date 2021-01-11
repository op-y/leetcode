package leetcode

import (
    "sort"
)

func findContentChildren(g []int, s []int) int {
    satisfied := 0
    sort.Ints(g)
    sort.Ints(s)
    i:=0
    j:=0
    for i<len(g) && j<len(s) {
        if s[j] >= g[i] {
            satisfied++
            i++
            j++
        } else {
            j++
        }
    }
    return satisfied
}
