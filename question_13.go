package leetcode

func romanToInt(s string) int {
    result := 0
    length := len(s)
    for i := 0; i < length; i++ {
        v := s[i]
        if v == 'M' {
            result += 1000
            continue
        }
        if v == 'D' {
            result += 500
            continue
        }
        if v == 'C' {
            if i+1 < length && s[i+1] == 'M' {
                result += 900
                i++
                continue
            }
            if i+1 < length && s[i+1] == 'D' {
                result += 400
                i++
                continue
            }
            result += 100
            continue
        }
        if v == 'L' {
            result += 50
            continue
        }
        if v == 'X' {
            if i+1 < length && s[i+1] == 'C' {
                result += 90
                i++
                continue
            }
            if i+1 < length && s[i+1] == 'L' {
                result += 40
                i++
                continue
            }
            result += 10
            continue
        }
        if v == 'V' {
            result += 5
            continue
        }
        if v == 'I' {
            if i+1 < length && s[i+1] == 'X' {
                result += 9
                i++
                continue
            }
            if i+1 < length && s[i+1] == 'V' {
                result += 4
                i++
                continue
            }
            result += 1
            continue
        }
    }
    return result
}
