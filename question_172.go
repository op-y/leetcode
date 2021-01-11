package leetcode

func trailingZeroes(n int) int {
    result := 0
    for factor:=5, i:=n/factor; i != 0; factor*=5 {
        result += i
    }
    return result
}
