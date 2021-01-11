package leetcode

func canConstruct(ransomNote string, magazine string) bool {
    if ransomNote == "" && magazine == "" {
        return true
    }
    if ransomNote != "" && magazine == "" {
        return false
    }

    dict := make(map[byte]int, 0)
    for _, b := range []byte(magazine) {
        if c, ok := dict[b]; ok {
            dict[b] = c + 1
        } else {
            dict[b] = 1
        }
    }

    for _, b := range []byte(ransomNote) {
        if c, ok := dict[b]; ok {
            if c == 0 {
                return false
            }
            dict[b] = c - 1
        } else {
            return false
        }
    }

    return true
}
