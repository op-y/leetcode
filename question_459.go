package leetcode

func repeatedSubstringPattern(s string) bool {
    if s == "" {
        return false
    }
    ss := []byte(s)
    length := len(ss)
    mid := length / 2
    for l:=1; l<=mid; l++ {
        if length % l != 0 {
            continue
        }
        sub := ss[0:l]
        ok := true
        for i:=l;i<length;i+=l {
            for j:=0;j<l;j++ {
                if sub[j] != ss[i+j] {
                    ok = false
                    break
                }
            }
            if !ok {
                break
            }
        }
        if ok {
            return true
        }
    }
    return false
}
