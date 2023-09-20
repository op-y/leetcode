package leetcode

import (
    "fmt"
)

func topKstrings( strings []string ,  k int ) [][]string {
    dict := make(map[string]int)
    for _, s := range strings {
        if c, ok := dict[s]; ok {
            dict[s] = c + 1
        } else {
            dict[s] = 1
        }
    }
    fmt.Printf("dict: %v \n", dict)
    ans := make([][]string, 0)
    list := make([]string, 0)
    for s, c := range dict {
        if len(list) < k {
            isInserted := false
            for i := 0; i < len(list); i++ {
                if dict[list[i]] >= c {
                    i++
                } else {
                    for j := len(list); j > i; j-- {
                        list[j] = list[j - 1]
                    }
                    list[i] = s
                    isInserted = true
                    break
                }
            }
            if !isInserted {
                list = append(list, s)
            }
            fmt.Printf("less then k, %d list: %v \n", len(list), list)
        } else {
            for i := 0; i < k; i++ {
                if dict[list[i]] >= c {
                    i++
                } else {
                    for j:=k-1; j > i; j-- {
                        list[j] = list[j - 1]
                    }
                    list[i] = s
                    break
                }
            }
            fmt.Printf("equal k, %d list: %v \n", len(list), list)
        }
    }
    for idx, s := range list {
        ans = append(ans, []string{fmt.Sprintf("%d", idx+1), s})
    }
    return ans
}
