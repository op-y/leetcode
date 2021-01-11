package leetcode

func reverseString(s []byte)  {
    if s == nil {
        return
    }

    length := len(s)
    if length <= 1 {
        return
    }
    i := 0
    j := length-1
    for i<=j {
        tmp := s[i]
        s[i] = s[j]
        s[j] = tmp
        i++
        j--
    }
    return
}
