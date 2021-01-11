package leetcode

func longestCommonPrefix(strs []string) string {
    strcount := len(strs)
    if strcount == 0 {
        return ""
    }
    if strcount == 1 {
        return strs[0]
    }
    diff := false
    idx := 0
    for {
        if len(strs[0]) <= idx {
            break
        }
        for i:=1; i<strcount; i++ {
            if len(strs[i]) <= idx || strs[i][idx] != strs[0][idx] {
                diff = true
                break
            }
        }
        if diff {
            break
        }
        idx++
    }
    return strs[0][0:idx]
}
