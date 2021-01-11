package leetcode

func isIsomorphic(s string, t string) bool {
    ss := []byte(s)
    st := []byte(t)
    if len(ss) != len(st) {
        return false
    }
    dict1 := make(map[byte]byte, 26)
    dict2 := make(map[byte]byte, 26)
    length := len(ss)
    for i:=0; i<length; i++ {
        if ch, ok := dict1[ss[i]]; ok {
            if ch != st[i] {
                return false
            }
        } else {
            if _, ok := dict2[st[i]]; ok {
                return false
            } else {
                dict1[ss[i]] = st[i]
                dict2[st[i]] = ss[i]
            }
        }
    }
    return true
}
