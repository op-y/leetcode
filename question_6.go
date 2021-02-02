package leetcode

func convert(s string, numRows int) string {
    if numRows <= 1 {
        return s
    }

    rows := make([][]byte, numRows)
    for i:=0; i<numRows; i++ {
        rows[i] = make([]byte, 0)
    }

    ss := []byte(s)
    up := true
    cursor := 0
    for _, c := range ss {
        rows[cursor] = append(rows[cursor], c)
        if up &&  cursor < numRows - 1 {
            cursor++
            continue
        }
        if up &&  cursor == numRows - 1 {
            up = false
            cursor--
            continue
        }
        if !up &&  cursor > 0 {
            cursor--
            continue
        }
        if !up &&  cursor == 0 {
            up = true
            cursor++
            continue
        }
    }

    ans := ""
    for i:=0; i < numRows; i++ {
        ans = ans + string(rows[i])
    }
    return ans
}
