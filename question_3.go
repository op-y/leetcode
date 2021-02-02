package leetcode

func lengthOfLongestSubstring(s string) int {
    if s == "" {
        return 0
    }

    ss := []byte(s)
    dict := make(map[byte]int, 0)

    max := 1

    start := 0
    end := 0
    current := 0
    for idx, c := range ss {
        if pos, ok := dict[c]; ok {
            i := start
            for i<=pos {
                delete(dict, ss[i])
                i++
            }
            if i >= len(ss) {
                start = i - 1
            } else {
                start = i
            }
            end = idx
            dict[c] = idx
            current = end - start + 1
            // 应该执行不到这个if
            if current > max {
                max = current
            }
        } else {
            end = idx
            dict[c] = idx
            current += 1
            if current > max {
                max = current
            }
        }
    }

    return max
}
