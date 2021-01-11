package leetcode

import (
    "fmt"
)

func convertToTitle(n int) string {
    result := ""
    digits := make([]int, 0)
    dict := map[int]string {
        1:"A",
        2:"B",
        3:"C",
        4:"D",
        5:"E",
        6:"F",
        7:"G",
        8:"H",
        9:"I",
        10:"J",
        11:"K",
        12:"L",
        13:"M",
        14:"N",
        15:"O",
        16:"P",
        17:"Q",
        18:"R",
        19:"S",
        20:"T",
        21:"U",
        22:"V",
        23:"W",
        24:"X",
        25:"Y",
        26:"Z"}
    for n > 26 {
        p := n / 26
        m := n % 26
        digits = append(digits, m)
        n = p
    }
    if n > 0 {
        digits = append(digits, n)
    }
    fmt.Printf("%v\n", digits)
    for i, d := range digits {
        if d <= 0 {
            if i == len(digits) - 1 {
                digits = digits[0:len(digits) - 1]
                break
            }
            fmt.Printf("[%d] %d carray\n", i, d)
            digits[i] += 26
            digits[i+1] -= 1
        }
    }
    fmt.Printf("%v\n", digits)
    for  _, d := range digits {
        result = dict[d] + result
    }
    return result
}
