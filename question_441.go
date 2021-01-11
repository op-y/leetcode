package leetcode

func arrangeCoins(n int) int {
    line := 0
    next := 1
    for n >= next {
        line++
        n -= next
        next++
    }
    return line
}
