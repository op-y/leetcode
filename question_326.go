package leetcode

func isPowerOfThree(n int) bool {
    if n <= 0 {
        return false
    }

    if n == 1 {
        return true
    }

    for n > 0 {
        if n % 3 != 0 {
            return false
        }
        n = n / 3
        if n == 1 {
            return true
        }
    }
    return true
}
