package leetcode

func isSubsequence(s string, t string) bool {
    if s == "" {
        return true
    }

    if t == "" {
        return false
    }

    ss := []byte(s)
    ts := []byte(t)
    lenss := len(ss)
    lenst := len(ts)
    if lenst < lenss {
        return false
    }

    i := 0
    j := 0
    for i<lenss && j<lenst {
        if ss[i] == ts[j] {
            i++
            j++
        } else {
            j++
        }
    }

    if i == lenss {
        return true
    }
    return false
}
