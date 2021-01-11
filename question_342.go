package leetcode

func isPowerOfFour(n int) bool {
    if n < 0 {
        return false
    }

    if n == 1 {
        return true
    }

    if n & (n-1) != 0 {
         return false
    }

    for n > 1 {
        n = n >> 2
    }
    if n == 1 {
        return true
    } else {
        return false
    }
}
