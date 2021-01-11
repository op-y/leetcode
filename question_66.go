package leetcode

func plusOne(digits []int) []int {
    if digits == nil || len(digits) == 0 {
        return []int{1}
    }
    carrier := 1
    sum := 0
    for i := len(digits) - 1; i>=0; i-- {
        sum = digits[i] + carrier
        if sum > 9 {
            carrier = 1
            digits[i] = sum - 10
        } else {
            carrier = 0
            digits[i] = sum
        }
    }
    if carrier > 0 {
        digits = append([]int{1}, digits...)
    }
    return digits
}
