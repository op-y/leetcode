package leetcode

func lengthOfLastWord(s string) int {
    if len(s) == 0 {
        return 0
    }
    start := false
    length := 0
    for _, ch := range s {
        if !start && ch != ' ' {
            start = true
            length = 1
        } else if start && ch != ' ' {
            length++
        } else if start && ch == ' ' {
            start = false
        } else {
            continue
        }
    }
    return length
}
