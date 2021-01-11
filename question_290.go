package leetcode

import (
    "strings"
)

func wordPattern(pattern string, s string) bool {
     dict1 := make(map[byte]string, 0)
     dict2 := make(map[string]byte, 0)
     ss := strings.Split(s , " ")
     sp := []byte(pattern)
     if len(ss) != len(sp) {
        return false
     }
     for i:=0; i<len(pattern); i++ {
        if v, ok := dict1[sp[i]]; ok {
            if v != ss[i] {
                return false
            }
        } else {
            dict1[sp[i]] = ss[i]
        }

        if v, ok := dict2[ss[i]]; ok{
            if v != sp[i] {
                return false
            }
        } else {
            dict2[ss[i]] = sp[i]
        }
     }
     return true
}
