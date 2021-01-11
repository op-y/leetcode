package leetcode

func addBinary(a string, b string) string {
    sum := ""
    carrier := false
    i, j :=len(a) - 1, len(b) - 1
    as := []byte(a)
    bs := []byte(b)
    for i>=0 && j>=0 {
        if as[i] == '0' && bs[j] == '0' {
            if carrier {
                sum = "1" + sum
                carrier = false
            } else {
                sum = "0" + sum
                carrier = false
            }
        }
        if as[i] == '1' && bs[j] == '0' {
            if carrier {
                sum = "0" + sum
                carrier = true
            } else {
                sum = "1" + sum
                carrier = false
            }
        }
        if as[i] == '0' && bs[j] == '1' {
            if carrier {
                sum = "0" + sum
                carrier = true
            } else {
                sum = "1" + sum
                carrier = false
            }
        }
        if as[i] == '1' && bs[j] == '1' {
            if carrier {
                sum = "1" + sum
                carrier = true
            } else {
                sum = "0" + sum
                carrier = true
            }
        }
        i--
        j--
    }
    for i >= 0 {
        if as[i] == '0' {
            if carrier {
                sum = "1" + sum
                carrier = false
            } else {
                sum = "0" + sum
                carrier = false
            }
        } else {
            if carrier {
                sum = "0" + sum
                carrier = true
            } else {
                sum = "1" + sum
                carrier = false
            }
        }
        i--
    }
    for j >=0 {
        if bs[j] == '0' {
            if carrier {
                sum = "1" + sum
                carrier = false
            } else {
                sum = "0" + sum
                carrier = false
            }
        } else {
            if carrier {
                sum = "0" + sum
                carrier = true
            } else {
                sum = "1" + sum
                carrier = false
            }
        }
        j--
    }
    if carrier {
        sum = "1" + sum
    }
    return sum
}
