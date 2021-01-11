package leetcode

func reverseVowels(s string) string {
    ss := []byte(s)

    if ss == nil {
        return
    }

    length := len(ss)
    if length <= 1 {
        return
    }

    dict := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

    i := 0
    j := length-1
    for i<=j {
        if _, ok := dict[i]; !ok{
            i++
            continue
        }
        if _, ok := dict[j]; !ok{
            j--
            continue
        }

        tmp := ss[i]
        ss[i] = ss[j]
        ss[j] = tmp
        i++
        j--
    }
    return string(ss)
}
