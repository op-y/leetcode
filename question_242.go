package leetcode

func isAnagram(s string, t string) bool {
    dict := make(map[byte]int, 0)
    ss := []byte(s)
    st := []byte(t)
    for _, b := range ss {
        if c, ok := dict[b]; ok {
            dict[b] = c + 1
        } else {
            dict[b] = 1
        }
    }
    for _, b := range st {
        if c, ok := dict[b]; ok {
            dict[b] = c - 1
        } else {
            return false
        }
    }
    for _, v := range dict {
        if v != 0 {
            return false
        }
    }
    return true
}
