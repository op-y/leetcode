package leetcode

func missingNumber(nums []int) int {
    n := len(nums)
    if n == 0 {
        return 0
    }
    expect := 0
    if n % 2 == 0 {
        expect = (n / 2) * (n + 1)
    } else {
        expect = ((n + 1) / 2) * n
    }
    sum := 0
    for _, v := range nums {
        sum += v
    }
    return expect - sum
}
