package leetcode

func isPerfectSquare(num int) bool {
    if num < 0 {
        return false
    }

    if num == 0 {
        return true
    }

    n := float64(num)
    x0 := n / float64(2)
    x1 := (x0 + n / x0) / float64(2)
    for x0 - x1 >= 1e-10  {
        x0 = x1
        x1 = (x0 + n / x0) / float64(2)
    }
    ans := int(x1)
    if ans * ans == num {
        return true
    }
    ans--
    if ans * ans == num {
        return true
    }
    return false
}
