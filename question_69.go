package leetcode

func mySqrt(x int) int {
    if x <= 0 {
        return 0
    }

    n := float64(x)
    x0 := n / float64(2)
    x1 := (x0 + n / x0) / float64(2)
    for x0 - x1 >= 1e-10  {
        x0 = x1
        x1 = (x0 + n / x0) / float64(2)
    }
    return int(x1)
}
