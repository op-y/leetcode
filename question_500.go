package leetcode

func findWords(words []string) []string {
    dict := map[byte]int{
       'q':1, 'Q':1, 'w':1, 'W':1, 'e':1, 'E':1, 'r':1, 'R':1, 't':1, 'T':1, 'y':1, 'Y':1, 'u':1, 'U':1, 'i':1, 'I':1, 'o':1, 'O':1, 'p':1, 'P':1,
       'a':2, 'A':2, 's':2, 'S':2, 'd':2, 'D':2, 'f':2, 'F':2, 'g':2, 'G':2, 'h':2, 'H':2, 'j':2, 'J':2, 'k':2, 'K':2, 'l':2, 'L':2,
       'z':3, 'Z':3, 'x':3, 'X':3, 'c':3, 'C':3, 'v':3, 'V':3, 'b':3, 'B':3, 'n':3, 'N':3, 'm':3, 'M':3}
    result := make([]string, 0)
    if words == nil {
        return nil
    }
    for _, s := range words {
        if s == "" {
            result = append(result, s)
            continue
        }
        ss := []byte(s)
        if len(ss) == 1 {
            result = append(result, s)
            continue
        }
        ok := true
        for i:=0; i<len(ss)-1; i++ {
            if dict[ss[i]] != dict[ss[i+1]] {
                ok = false
                break
            }
        }
        if ok {
            result = append(result, s)
        }
    }
    return result
}
