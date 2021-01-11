package leetcode

func firstUniqChar(s string) int {
    if s == "" {
        return -1
    }

    dict := make(map[byte][]int, 0)
    ss := []byte(s)
    length := len(ss)
    if length == 1 {
        return 0
    }

    for i, b := range ss {
        if c, ok := dict[b]; ok {
            c[1] = c[1] + 1
        } else {
            dict[b] = []int{i, 1}
        }
    }

    idx := length
    for _, v := range dict {
        if v[1] == 1 {
            if v[0] < idx {
                idx = v[0]
            }
        }
    }

    if idx == length {
        return -1
    }
    return idx
}
