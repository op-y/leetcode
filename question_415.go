package leetcode

func addStrings(num1 string, num2 string) string {
    if num1 == "" {
        return num2
    }
    if num2 == "" {
        return num1
    }

    sum := make([]byte, 0)
    snum1 := []byte(num1)
    snum2 := []byte(num2)
    len1 := len(snum1)
    len2 := len(snum2)

    i := len1 - 1
    j := len2 - 1

    carrier := false

    for i>=0 && j>=0 {
        v := int(snum1[i] - '0') + int(snum2[j] - '0')
        if carrier {
            v += 1
        }

        if v >= 10 {
           v = v - 10
           carrier = true
        } else {
           carrier = false
        }

        b := byte(v + 48)
        sum = append([]byte{b}, sum...)

        i--
        j--
    }

    for i>=0 {
        v := int(snum1[i] - '0')
        if carrier {
            v += 1
        }

        if v >= 10 {
           v = v - 10
           carrier = true
        } else {
           carrier = false
        }

        b := byte(v + 48)
        sum = append([]byte{b}, sum...)

        i--
    }

    for j>=0 {
        v := int(snum2[j] - '0')
        if carrier {
            v += 1
        }

        if v >= 10 {
           v = v - 10
           carrier = true
        } else {
           carrier = false
        }

        b := byte(v + 48)
        sum = append([]byte{b}, sum...)

        j--
    }

    if carrier {
        sum = append([]byte{'1'}, sum...)
    }

    return string(sum)
}
